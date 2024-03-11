package controller

import (
	"github.com/gin-gonic/gin"
)

// Download 下载文件 请求路径：127.0.0.1:9888/api/test/download
func Download(ctx *gin.Context) {
	fileName := "file.txt"
	filePath := "download/" + fileName
	ctx.Header("Content-Disposition", "attachment;filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(filePath)
}
