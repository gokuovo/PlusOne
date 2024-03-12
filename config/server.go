package config

import (
	"PlusOne/middleware"
	"PlusOne/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Server(port int64) {
	// 创建Gin引擎
	router := gin.Default()

	// 设置日志输出到文件
	middleware.SetupLogger()

	// 注册日志中间件
	router.Use(middleware.LogMiddleware)

	//注册路由
	route.Route(router)

	println(fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, 31, msg, 0x1B))

	// 启动Gin应用
	router.Run(fmt.Sprintf(":%s", strconv.FormatInt(port, 10)))
}
