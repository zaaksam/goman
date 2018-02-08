package config

// RunModeType 运行模式
type RunModeType int

const (
	RUN_MODE_WEB = iota
	RUN_MODE_APP
	RUN_MODE_DOCKER
)
