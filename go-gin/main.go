package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return new(Controller)
}

type TBindParameter struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func (c *Controller) TBind(ctx *gin.Context) {
	var param TBindParameter
	if err := ctx.Bind(&param); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"code": 500, "msg": err})
		return
	}
	// if err := ctx.ShouldBind(&param); err != nil {
	// 	ctx.JSON(http.StatusBadGateway, gin.H{"code": 500, "msg": err})
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "", "data": param})
}

func main() {
	route := gin.Default()
	route.GET("list", NewController().TBind)
	route.POST("list", NewController().TBind)
	if err := route.Run(":9888"); err != nil {
		log.Printf("服务器启动失败:%v\n", err)
	}
}
