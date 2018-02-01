package model

import (
	"net/http"
)

type RequestModel struct {
	N       int         `json:"n"`
	C       int         `json:"c"`
	Timeout int         `json:"timeout"`
	Method  string      `json:"method"`
	URL     string      `json:"url"`
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
}
