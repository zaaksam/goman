package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/imroc/req"
	"github.com/zaaksam/goman/go/model"
)

// Req Http请求对象
var Req reqService

type reqService struct{}

func (s *reqService) DoRequest() (resp *req.Resp, err error) {
	r := req.New()

	resp, err = r.Do("", "")
	return
}

func (s *reqService) Do(req *model.RequestModel) (res *model.ResponseModel, err error) {
	httpClient := http.Client{}
	var httpReq *http.Request

	url := req.URL
	if strings.Index(url, "://") == -1 {
		url = "http://" + url
	}

	// if len(req.Bodys) > 0 {
	// 	req.Body = ""

	// 	c := 0
	// 	for k, v := range req.Bodys {
	// 		if c > 0 {
	// 			req.Body += "&"
	// 		}

	// 		for i, l := 0, len(v); i < l; i++ {
	// 			if i > 0 {
	// 				req.Body += "&"
	// 			}

	// 			req.Body += k + "=" + netURL.QueryEscape(v[i])
	// 		}

	// 		c++
	// 	}
	// }

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

	//添加Cookie
	// for i := range req.Cookies {
	// 	httpReq.AddCookie(req.Cookies[i])
	// }

	tn := time.Now()
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return
	}

	duration := fmt.Sprintf("%s", time.Now().Sub(tn))

	//读取响应结果
	defer httpRes.Body.Close()
	var resBody []byte
	resBody, err = ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return
	}

	var resCode []byte
	resCode, err = httputil.DumpResponse(httpRes, false)
	if err != nil {
		return
	}

	res = &model.ResponseModel{
		Status:     httpRes.Status,
		StatusCode: httpRes.StatusCode,
		Duration:   duration,
		Proto:      httpRes.Proto,
		Headers:    httpRes.Header,
		Cookies:    httpRes.Cookies(),
		Body:       string(resBody),
		Code:       string(resCode),
	}
	return
}
