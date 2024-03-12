package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type viewController struct {
}

var _viewController *viewController

func GetViewController() *viewController {
	return _viewController
}

func (v *viewController) View(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "hello",
		"content": "内容",
	})
}
