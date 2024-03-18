package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

func ApiRouter(r gin.IRouter) {
	api := IRouter{IRouter: r.Group("/api")}
	{
		api := api.Group("/download")
		api.GetPost("file", controller.GetDownLoadController().Download)
	}

	{
		api := api.Group("/bind")
		api.GetPost("bind", controller.GetBindController().Bind)
		api.GetPost("ShouldBind", controller.GetBindController().ShouldBind)
		api.GetPost("ValidatorShouldBind", controller.GetBindController().ValidatorShouldBind)
		api.GetPost("BindValidator", controller.GetBindController().BindValidator)
	}

	{
		api := api.Group("view")
		api.GetPost("details", controller.GetViewController().View)
	}
}
