package postgre

import (
	"github.com/gogf/gf/os/glog"
	"os/exec"
	"postgre_mate_go/pkg/path"
)

func Start() {
	startPostgre()
}

func startPostgre() {
	command := exec.Command("/bin/bash", path.PostgreStartScript)
	err := command.Start()
	if err != nil {
		glog.Error("start postgre server failed")
	}
	command.Wait()
}
