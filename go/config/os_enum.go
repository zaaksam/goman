package config

// OSType 运行系统
type OSType int

func (os OSType) String() string {
	s := AppConf.goos
	if os == OS_TYPE_UNSUPPORT {
		s = "不支持的系统：" + s
	}

	return s
}

const (
	OS_TYPE_UNSUPPORT = iota
	OS_TYPE_WINDOWS
	OS_TYPE_DARWIN
	OS_TYPE_LINUX
)
