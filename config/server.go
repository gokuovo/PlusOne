package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func server(port int64) {
	// 创建Gin引擎
	router := gin.Default()

	// 定义路由和处理函数
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})
	str := strconv.FormatInt(port, 10)
	// 启动Gin应用
	router.Run(fmt.Sprintf(":%s", str))
}
