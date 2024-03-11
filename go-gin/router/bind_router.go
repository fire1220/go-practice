package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

func BindRouter(r gin.IRouter) {
	api := IRouter{IRouter: r.Group("/api")}
	{
		api := api.Group("/test")
		api.GetPost("bind", controller.Bind)
		api.GetPost("ShouldBind", controller.ShouldBind)
	}
}
