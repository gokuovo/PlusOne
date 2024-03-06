package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"io"
	"log"
	"os"
	"time"
)

func SetupLogger() {
	// 创建一个新的日志文件，并设置相关参数
	logFile := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    5,   // 每个日志文件的最大尺寸，单位是MB
		MaxBackups: 365, // 保留的旧日志文件最大数量
		MaxAge:     365, // 保留的旧日志文件最大天数
	}

	// 创建一个新的多写器，同时输出到终端和日志文件
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 设置Gin框架的默认日志输出到多写器
	gin.DefaultWriter = multiWriter
	log.SetOutput(multiWriter)
}

func LogMiddleware(c *gin.Context) {
	start := time.Now()

	// 在处理请求之前打印日志
	log.Printf("Started %s %s", c.Request.Method, c.Request.URL.Path)

	// 调用下一个处理程序
	c.Next()

	// 在处理请求之后打印日志
	duration := time.Since(start)
	log.Printf("Completed %s %s in %s", c.Request.Method, c.Request.URL.Path, duration)
}
