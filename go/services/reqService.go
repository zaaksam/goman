package services

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"strings"
	"sync"
	"time"

	"github.com/zaaksam/goman/go/model"
)

// Req Http请求对象
var Req reqService

type reqService struct{}

func (s *reqService) Do(req *model.RequestModel) (list []*model.ResponseModel, err error) {
	var httpReq *http.Request

	url := req.URL
	if strings.Index(url, "://") == -1 {
		url = "http://" + url
	}

	method := strings.ToUpper(req.Method)
	if req.Body != "" {
		httpReq, err = http.NewRequest(method, url, strings.NewReader(req.Body))
	} else {
		httpReq, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return
	}

	//重新组装Header，使其标准化
	for k, v := range req.Headers {
		for i, l := 0, len(v); i < l; i++ {
			httpReq.Header.Add(k, v[i])
		}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConnsPerHost: s.min(req.C, 500),
		DisableCompression:  false,
		DisableKeepAlives:   false,
		// Proxy:               http.ProxyURL(b.ProxyAddr),
	}
	// if b.H2 {
	// 	http2.ConfigureTransport(tr)
	// } else {
	tr.TLSNextProto = make(map[string]func(string, *tls.Conn) http.RoundTripper)
	// }
	httpClient := &http.Client{Transport: tr}

	if req.Timeout > 0 {
		httpClient.Timeout = time.Duration(req.Timeout) * time.Second
	}

	//请求数少于并发数时，并发数调整为请求数一致
	c := req.C
	if req.N < c {
		c = req.N
	}

	z := req.N / c
	y := req.N % c
	if y > 0 {
		//当请求数取余并发数时大于0时，将每协程处理请求数字加1
		z++
	}

	doneChan := make(chan struct{})
	resChan := make(chan *model.ResponseModel, c*1000)
	list = make([]*model.ResponseModel, 0, req.N)

	var wg sync.WaitGroup
	wg.Add(c)

	//并发请求协程
	for i := 0; i < c; i++ {
		l := z
		if y > 0 && (i+1) > y {
			//对余数请求的协程处理上限减1
			l--
		}

		go func(ll int) {
			for j := 0; j < ll; j++ {
				s.do(httpClient, httpReq, resChan)
			}
			wg.Done()
		}(l)
	}

	go func() {
		for res := range resChan {
			list = append(list, res)
		}
		close(doneChan)
	}()

	wg.Wait()
	close(resChan)

	<-doneChan
	return
}

func (s *reqService) do(c *http.Client, r *http.Request, resChan chan *model.ResponseModel) {
	res := &model.ResponseModel{}
	start := time.Now()
	var dnsStart, connStart, resStart, reqStart, delayStart time.Time
	var dnsDuration, connDuration, resDuration, reqDuration, delayDuration time.Duration
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			dnsDuration = time.Now().Sub(dnsStart)
		},
		GetConn: func(h string) {
			connStart = time.Now()
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			if !connInfo.Reused {
				connDuration = time.Now().Sub(connStart)
			}
			reqStart = time.Now()
		},
		WroteRequest: func(w httptrace.WroteRequestInfo) {
			reqDuration = time.Now().Sub(reqStart)
			delayStart = time.Now()
		},
		GotFirstResponseByte: func() {
			delayDuration = time.Now().Sub(delayStart)
			resStart = time.Now()
		},
	}

	r = s.cloneRequest(r)
	r = r.WithContext(httptrace.WithClientTrace(r.Context(), trace))

	httpRes, err := c.Do(r)
	if err != nil {
		res.Error = err.Error()
		resChan <- res
		return
	}

	t := time.Now()
	resDuration = t.Sub(resStart)
	finishDuration := t.Sub(start)

	var resBody []byte
	resBody, err = ioutil.ReadAll(httpRes.Body)
	httpRes.Body.Close()
	if err != nil {
		res.Error = err.Error()
		resChan <- res
		return
	}

	var resCode []byte
	resCode, err = httputil.DumpResponse(httpRes, false)
	if err != nil {
		res.Error = err.Error()
		resChan <- res
		return
	}

	res.Status = httpRes.Status
	res.StatusCode = httpRes.StatusCode
	res.ContentLength = httpRes.ContentLength
	res.Proto = httpRes.Proto
	res.Headers = httpRes.Header
	res.Cookies = httpRes.Cookies()
	res.Body = string(resBody)
	res.Code = string(resCode)

	if res.ContentLength < 0 {
		res.ContentLength = int64(len(res.Body))
	}

	res.Duration.DNS = s.formatDuration(dnsDuration)
	res.Duration.Conn = s.formatDuration(connDuration)
	res.Duration.Req = s.formatDuration(reqDuration)
	res.Duration.Res = s.formatDuration(resDuration)
	res.Duration.Delay = s.formatDuration(delayDuration)
	res.Duration.Finish = s.formatDuration(finishDuration)

	resChan <- res
}

func (s *reqService) cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r

	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}

	if r.GetBody != nil {
		r2.Body, _ = r.GetBody()
	}

	return r2
}

func (s *reqService) formatDuration(d time.Duration) string {
	ns := d.Nanoseconds()
	if ns < 10 {
		return fmt.Sprintf("%dns", ns)
	}

	us := float64(ns) / 1000
	if us < 10 {
		return fmt.Sprintf("%.2fµs", us)
	}

	ms := us / 1000
	if ms < 1000 {
		return fmt.Sprintf("%.2fms", ms)
	}

	return fmt.Sprintf("%.2fs", ms/1000)
}

func (s *reqService) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
