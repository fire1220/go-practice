package controller

import (
	"github.com/gin-gonic/gin"
	"main/ginvalidate"
	"net/http"
)

type BaseController struct {
}

func (b *BaseController) Validate(ctx *gin.Context, param any) bool {
	if ok, errs := ginvalidate.SimpleValidate(ctx, param); !ok {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"code": 405, "msg": errs[0].Error()})
		return false
	}
	return true
}
