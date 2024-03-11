package middleware

import (
	"runtime/debug"

	"codeup.aliyun.com/xhey/server/serverkit/v2/logger"
	"codeup.aliyun.com/xhey/server/serverkit/v2/rest"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				connId, _ := c.Get("Conn-ID")
				logger.Error("[%v] error: %s, stack: %s", connId, err, string(debug.Stack()))
				rest.DoResponse(c, rest.StatusInternalServerError, nil)
			}
		}()
		c.Next()
	}
}
