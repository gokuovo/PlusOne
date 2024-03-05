package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建Gin引擎
	router := gin.Default()

	// 定义路由和处理函数
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 启动Gin应用
	router.Run(":8080")
}
