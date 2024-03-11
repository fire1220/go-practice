package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

func DownloadRouter(r gin.IRouter) {
	api := IRouter{IRouter: r.Group("/api")}
	{
		api := api.Group("/test")
		api.GetPost("download", controller.Download)
	}
}
