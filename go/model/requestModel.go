package model

import (
	"net/http"
)

type RequestModel struct {
	Method  string      `json:"method"`
	URL     string      `json:"url"`
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
}
