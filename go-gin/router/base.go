package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

type IRouter struct {
	gin.IRouter
}

func (i IRouter) GetPost(s string, fn ...gin.HandlerFunc) {
	i.GET(s, fn...)
	i.POST(s, fn...)
}

func (i IRouter) Group(s string, fn ...gin.HandlerFunc) IRouter {
	return IRouter{IRouter: i.IRouter.Group(s, fn...)}
}

func RegRouter(r gin.IRouter) {
	api := IRouter{IRouter: r.Group("/api")}
	{
		api := api.Group("/test")
		api.GetPost("list", controller.NewController().TBind)
	}
}
