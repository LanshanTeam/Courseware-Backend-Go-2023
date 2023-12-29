package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个Gin路由器
	router := gin.Default()

	// 创建一个路由处理器
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 启动HTTP服务器并监听端口
	router.Run(":8080")
}
