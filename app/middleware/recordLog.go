package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func RecordPostLog() gin.HandlerFunc {
	// 打开文件准备写入日志
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	logFile, err := os.OpenFile(path+"/../logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// 设置输出目标为文件
	log.SetOutput(logFile)

	return func(context *gin.Context) {
		startTime := time.Now()

		// 处理请求
		context.Next()

		// 记录结束时间
		endTime := time.Now()

		// 计算请求耗时
		latency := endTime.Sub(startTime)

		// 获取请求信息
		requestMethod := context.Request.Method
		requestURL := context.Request.URL.String()
		requestProto := context.Request.Proto

		// 获取响应信息
		responseStatus := context.Writer.Status()
		responseSize := context.Writer.Size()

		// 检查是否有错误发生
		errs := context.Errors.ByType(gin.ErrorTypeAny)
		if len(errs) > 0 {
			// 将错误记录到Gin框架的错误日志中
			log.Printf("[%v] %v %v %v|respSize:%v|%v - Errors: %v\n", requestMethod, requestURL, requestProto, responseStatus, responseSize, latency, errs)
		} else {
			// 记录请求日志
			log.Printf("[%v] %v %v %v|respSize:%v|%v\n", requestMethod, requestURL, requestProto, responseStatus, responseSize, latency)
		}
	}
}
