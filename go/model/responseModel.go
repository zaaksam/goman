package model

import (
	"net/http"
)

type ResponseModel struct {
	Status     string         `json:"status"`
	StatusCode int            `json:"statusCode"`
	Duration   string         `json:"duration"`
	Proto      string         `json:"proto"`
	Headers    http.Header    `json:"headers"`
	Cookies    []*http.Cookie `json:"cookies"`
	Body       string         `json:"body"`
	Code       string         `json:"code"`
}
