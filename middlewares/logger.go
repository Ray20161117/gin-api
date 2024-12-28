/**
 * 日志log中间件
 */
package middlewares

import (
	"bytes"
	"gin-api/pkg/log"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 保存请求体以便后续读取
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Error(err) // 记录错误
			return
		}
		bodyReader := io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Request.Body = bodyReader

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 计算执行时间
		latencyTime := endTime.Sub(startTime).Milliseconds()

		// 请求方式和URI
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI

		// 请求头和协议
		header := c.Request.Header
		proto := c.Request.Proto

		// 状态码和客户端IP
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// 获取最后一个错误，如果有错误的话
		lastError := c.Errors.Last()

		// 记录日志
		logger := log.Log()
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"header":       header,
			"proto":        proto,
			"body":         bodyBytes, // 使用保存的请求体
			"err":          lastError,
		}).Info("Request processed")

		// 关闭请求体读取器
		if err := bodyReader.Close(); err != nil {
			logger.WithError(err).Error("Failed to close body reader")
		}
	}
}
