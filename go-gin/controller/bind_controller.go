package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type bindController struct {
	*baseController
}

var _bindController = new(bindController)

func GetBindController() *bindController {
	return _bindController
}

type TBindParameter struct {
	Id   int    `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name"`
	Date string `json:"date" form:"date" binding:"datetime=2006-01"`
}

// Bind 接收参数(入股保存会相应更改响应码，RESTful风格，此时家变ctx.JSON修改相应码也是无效的)
func (b *bindController) Bind(ctx *gin.Context) {
	var param TBindParameter
	if err := ctx.Bind(&param); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"code": 500, "msg": err})
		return
	}
	// 修改响应码无效
	ctx.JSON(0, gin.H{"code": 200, "msg": "", "data": param})
}

// ShouldBind 接收参数(自定义修改响应码),接收参数失败不影响响应码
func (b *bindController) ShouldBind(ctx *gin.Context) {
	var param TBindParameter
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"code": 500, "msg": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "", "data": param})
}

// ValidatorShouldBind 验证器
func (b *bindController) ValidatorShouldBind(ctx *gin.Context) {
	var param TBindParameter
	if err := ctx.ShouldBind(&param); err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadGateway, gin.H{"code": 500, "msg": e.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "", "data": param})
}

// BindValidator 验证器(中文)
func (b *bindController) BindValidator(ctx *gin.Context) {
	var param TBindParameter
	if ok := b.Validate(ctx, &param); !ok {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "", "data": param})
}
