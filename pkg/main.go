package main

// driver
import (
	_ "github.com/lib/pq"
)

// route
import (
	_ "postgre_mate_go/pkg/router"
)

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"postgre_mate_go/pkg/config"
	"postgre_mate_go/pkg/postgre"
)

func main() {
	glog.Debug("This is a debug msg")
	glog.Info("This is a info msg")
	glog.Error("This is a error msg")
	if !config.RemoteMode {
		postgre.Start()
	}
	g.Server().Run()
}
