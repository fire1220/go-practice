package controller

import (
	"github.com/gin-gonic/gin"
	"main/ginvalidate"
	"net/http"
)

type baseController struct {
}

func (b *baseController) Validate(ctx *gin.Context, param any) bool {
	if ok, errs := ginvalidate.SimpleValidate(ctx, &param); !ok {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"code": 405, "msg": errs[0].Error()})
		return false
	}
	return true
}
