package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllowCrossDomain(c *gin.Context) {
	r := c.Request
	if r.RequestURI == "/" {
		return
	}

	w := c.Writer

	//设置Credentials为true后，Origin不允许为*，读取request中的Origin允许访问所有域
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	//允许真实请求中的自定义头
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "web-version")
	w.Header().Add("Access-Control-Allow-Headers", "x-user-id")
	w.Header().Add("Access-Control-Allow-Headers", "device_id")
	//允许真实请求的方法
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, HEAD, OPTIONS")
	//允许携带cookie等
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	//预请求缓存时间 20天
	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Content-Type", "application/json")

	method := c.Request.Method
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func CrossDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		AllowCrossDomain(c)
		c.Next()
	}
}
