package router

import (
	"github.com/gin-gonic/gin"
)

type IRouter struct {
	gin.IRouter
}

func (i IRouter) GetPost(s string, fn ...gin.HandlerFunc) {
	i.GET(s, fn...)
	i.POST(s, fn...)
}

func (i IRouter) Group(s string, fn ...gin.HandlerFunc) IRouter {
	return IRouter{IRouter: i.IRouter.Group(s, fn...)}
}

// RegRouter 注册路由
func RegRouter(r gin.IRouter) {
	ApiRouter(r)
}
