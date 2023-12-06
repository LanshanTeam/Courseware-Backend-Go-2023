package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//在这里使用了三个中间件
	r.Use(mid1(), mid2(), mid3())
	r.GET("/abc", func(context *gin.Context) {
		fmt.Println("1")
	})
	_ = r.Run(":8081")
}

func mid1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mid1 before")
		//先执行next,后执行after
		c.Next()
		fmt.Println("mid1 after")
	}
}
func mid2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mid2 before")
		c.Abort()
		//c.Next()
		fmt.Println("mid2 after")
	}
}

func mid3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mid3 before")
		c.Next()
		fmt.Println("mid3 after")
	}
}
