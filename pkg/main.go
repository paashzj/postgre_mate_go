package main

import (
	"github.com/gogf/gf/os/glog"
	"postgre_mate_go/pkg/postgre"
)

func main() {
	glog.Debug("This is a debug msg")
	glog.Info("This is a info msg")
	glog.Error("This is a error msg")
	postgre.Start()
}
