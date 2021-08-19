package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"net/http"
	"postgre_mate_go/pkg/dao"
)

func CreateHealth(r *ghttp.Request) {
	err := dao.InsertHealth()
	if err != nil {
		glog.Info("err occurred ", err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	r.Response.WriteStatusExit(http.StatusCreated)
}

func UpdateHealth(r *ghttp.Request) {
	err := dao.UpdateHealth()
	if err != nil {
		glog.Info("err occurred ", err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	r.Response.WriteStatusExit(http.StatusOK)
}

func DeleteHealth(r *ghttp.Request) {
	err := dao.DeleteHealth()
	if err != nil {
		glog.Info("err occurred ", err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	r.Response.WriteStatusExit(http.StatusNoContent)
}
