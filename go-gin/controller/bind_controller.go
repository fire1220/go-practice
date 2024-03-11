package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TBindParameter struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

// Bind 接收参数(入股保存会相应更改响应码，RESTful风格，此时家变ctx.JSON修改相应码也是无效的)
func Bind(ctx *gin.Context) {
	var param TBindParameter
	if err := ctx.Bind(&param); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"code": 500, "msg": err})
		return
	}
	// 修改响应码无效
	ctx.JSON(0, gin.H{"code": 200, "msg": "", "data": param})
}

// ShouldBind 接收参数(自定义修改响应码),接收参数失败不影响响应码
func ShouldBind(ctx *gin.Context) {
	var param TBindParameter
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"code": 500, "msg": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "", "data": param})
}
