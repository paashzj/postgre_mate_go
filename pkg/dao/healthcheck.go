package dao

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"postgre_mate_go/pkg/config"
	"postgre_mate_go/pkg/module"
	"strconv"
)

const (
	dbType      = "pgsql"
	healthDb    = "ttbb"
	healthTable = "health_check"
)

func init() {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host: config.Hostname,
				Port: strconv.Itoa(config.Port),
				User: config.Username,
				Pass: config.Password,
				Type: dbType,
			},
		},
	})
}

func InsertHealth() error {
	table := g.DB().Schema(healthDb).Table(healthTable)
	healthCheck := &module.HealthCheck{}
	healthCheck.Id = config.HealthCheckId
	s := "null string"
	healthCheck.StrNull = &s
	healthCheck.StrNotNull = "not null string"
	i := 5
	healthCheck.IntNull = &i
	healthCheck.IntNotNull = 5
	result, err := table.Insert(healthCheck)
	glog.Info(result)
	return err
}

func UpdateHealth() error {
	table := g.DB().Schema(healthDb).Table(healthTable).OmitEmpty()
	healthCheck := &module.HealthCheck{}
	healthCheck.Id = config.HealthCheckId
	table.Data(healthCheck).Where("id", config.HealthCheckId)
	result, err := table.Update()
	glog.Info(result)
	return err
}

func DeleteHealth() error {
	result, err := g.DB().Schema(healthDb).Table(healthTable).Where("id", config.HealthCheckId).Delete()
	glog.Info(result)
	return err
}
