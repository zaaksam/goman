package services

import (
	"time"
)

// App 应用对象
var App appService

type appService struct {
	hearbeat int64
}

func (s *appService) IsConnect() bool {
	isConnect := false

	diff := time.Now().UnixNano() - s.hearbeat
	if diff < 2*1000000000 {
		isConnect = true
	}

	return isConnect
}

func (s *appService) Heartbeat() {
	s.hearbeat = time.Now().UnixNano()
}
