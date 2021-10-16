package postgre

import (
	"github.com/gogf/gf/os/glog"
	"github.com/paashzj/gutil"
	"postgre_mate_go/pkg/path"
)

func Start() {
	startPostgre()
}

func startPostgre() {
	stdout, stderr, err := gutil.CallScript(path.PostgreStartScript)
	if err != nil {
		glog.Error("start postgre server failed")
	}
	glog.Info("std out is ", stdout)
	glog.Error("std err is ", stderr)
}
