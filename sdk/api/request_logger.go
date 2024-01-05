package api

import (
	"github.com/go-sanpang/go-admin-core/helper"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/go-sanpang/go-admin-core/debug/logger"
	"github.com/go-sanpang/go-admin-core/sdk"
)

type loggerKey struct{}

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *logger.Helper {
	var log *logger.Helper
	l, ok := c.Get(helper.LoggerKey)
	if ok {
		ok = false
		log, ok = l.(*logger.Helper)
		if ok {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := helper.GenerateMsgIDFromContext(c)
	log = logger.NewHelper(sdk.Runtime.GetLogger()).WithFields(map[string]interface{}{
		strings.ToLower(helper.TrafficKey): requestId,
	})
	return log
}

// SetRequestLogger 设置logger中间件
func SetRequestLogger(c *gin.Context) {
	requestId := helper.GenerateMsgIDFromContext(c)
	log := logger.NewHelper(sdk.Runtime.GetLogger()).WithFields(map[string]interface{}{
		strings.ToLower(helper.TrafficKey): requestId,
	})
	c.Set(helper.LoggerKey, log)
}
