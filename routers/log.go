/*
Author:hujun@ztgame.com
Date:  2020-12-22
*/
package routers

import (
	"github.com/gin-gonic/gin"
	"10.254.25.25:8080/hujun/zt_cmdb/utils"
	"10.254.25.25:8080/hujun/logger"
	"time"
)

func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		statusColor := logger.ColorForStatus(statusCode)
		methodColor := logger.ColorForMethod(method)
		resetColor := logger.Reset
		traceId := c.Request.Header.Get("trace_id")
		userEmail := c.Request.Header.Get("user_email")
		requestData := utils.GetRequestData(c)
		logger.Info("[GIN] %s%s%s %s%s %s%d%s cost:%.03f [ip:%s] [trace_id:%s] [user_email:%s]",
			methodColor, method, resetColor,
			c.Request.Host,
			utils.Cuts(requestData, utils.MaxLogDataLen),
			statusColor, statusCode, resetColor,
			end.Sub(start).Seconds(),
			clientIP, traceId, userEmail)
	}
}
