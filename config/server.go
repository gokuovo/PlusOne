package config

import (
	"PlusOne/middleware"
	"PlusOne/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func server(port int64) {
	// 创建Gin引擎
	router := gin.Default()

	// 设置日志输出到文件
	middleware.SetupLogger()

	// 注册日志中间件
	router.Use(middleware.LogMiddleware)
	// 定义路由和处理函数
	router.GET("/", func(c *gin.Context) {
		//var a []models.A
		a := make([]models.A, 0)
		DB.Find(&a)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 启动Gin应用
	router.Run(fmt.Sprintf(":%s", strconv.FormatInt(port, 10)))
}
