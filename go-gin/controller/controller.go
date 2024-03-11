package controller

import (
	"github.com/gin-gonic/gin"
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

// TBind 接收参数
func (c *Controller) TBind(ctx *gin.Context) {
	var param TBindParameter
	// if err := ctx.Bind(&param); err != nil {
	// 	ctx.JSON(http.StatusConflict, gin.H{"code": 500, "msg": err})
	// 	return
	// }
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"code": 500, "msg": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "", "data": param})
}

// Download 下载文件
func (c *Controller) Download(ctx *gin.Context) {
	fileName := "file.txt"
	filePath := "download/" + fileName
	ctx.Header("Content-Disposition", "attachment;filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(filePath)
}
