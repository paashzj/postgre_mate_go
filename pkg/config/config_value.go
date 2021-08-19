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
	Hostname      string
	Port          int
	Password      string
	Username      string
	HealthCheckId string
)

func init() {
	Cluster = gutil.GetEnvStr("CLUSTER", "default")
	LogFile = gutil.GetEnvStr("LOG_FILE", "")
	ConsoleLog = gutil.GetEnvBool("CONSOLE_LOG", true)
	ListenAddr = gutil.GetEnvStr("LISTEN_ADDR", "0.0.0.0")
	RemoteMode = gutil.GetEnvBool("REMOTE_MODE", true)
	Hostname = gutil.GetEnvStr("HOST", "127.0.0.1")
	Port = gutil.GetEnvInt("PORT", 5432)
	Username = gutil.GetEnvStr("USERNAME", "sh")
	Password = gutil.GetEnvStr("PASSWORD", "ttlovezj")
	HealthCheckId = gutil.GetEnvStr("HOSTNAME", "default")
}
