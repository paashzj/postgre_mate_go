package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"postgre_mate_go/pkg/api"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/v1/postgre/", func(group *ghttp.RouterGroup) {
		group.Group("/health", func(group *ghttp.RouterGroup) {
			group.POST("/", api.CreateHealth)
			group.PUT("/", api.UpdateHealth)
			group.DELETE("/", api.DeleteHealth)
		})
	})
}
