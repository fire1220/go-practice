package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/router"
)

func main() {
	g := gin.Default()
	router.RegRouter(g)
	if err := g.Run(":9888"); err != nil {
		log.Printf("服务器启动失败:%v\n", err)
	}
}
