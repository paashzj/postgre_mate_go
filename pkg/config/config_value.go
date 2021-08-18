package config

import (
	"github.com/paashzj/gutil"
)

// mate config
var (
	Cluster    string
	LogFile    string
	ConsoleLog bool
	ListenAddr string
	RemoteMode bool
)

// postgre config
var (
	Password string
	Hostname string
	Username string
)

func init() {
	Cluster = gutil.GetEnvStr("CLUSTER", "default")
	LogFile = gutil.GetEnvStr("LOG_FILE", "")
	ConsoleLog = gutil.GetEnvBool("CONSOLE_LOG", true)
	ListenAddr = gutil.GetEnvStr("LISTEN_ADDR", "0.0.0.0")
	RemoteMode = gutil.GetEnvBool("REMOTE_MODE", true)
	Username = gutil.GetEnvStr("USERNAME", "hzj")
	Password = gutil.GetEnvStr("PASSWORD", "Mysql@123")
	Hostname = gutil.GetEnvStr("HOSTNAME", "127.0.0.1:3306")
}
