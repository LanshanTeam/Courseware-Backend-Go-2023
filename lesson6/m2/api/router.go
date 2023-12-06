package api

import (
	"Courseware-Backend-Go-2023/lesson6/m2/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", register) // 注册
	r.POST("/login", login)       // 登录

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8088") // 跑在 8088 端口上
}
