package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"codeup.aliyun.com/xhey/server/serverkit/v2/logger"
	"codeup.aliyun.com/xhey/server/serverkit/v2/rest"

	"github.com/gin-gonic/gin"
)

var (
	connId uint64
)

var apiLogBlackList = map[string]interface{}{
	"/":        1,
	"/metrics": 1,
}

func RequestInLog(c *gin.Context) {
	r := c.Request
	if _, ok := apiLogBlackList[r.RequestURI]; ok {
		return
	}

	id := atomic.AddUint64(&connId, 1)
	c.Set("Conn-ID", id)
	r.Header.Set("Conn-ID", fmt.Sprintf("%v", id))

	startExecTime := time.Now()
	c.Set("startExecTime", startExecTime)

	if rest.IsEncrypted && strings.Contains(c.Request.RequestURI, "/next/") {
		return
	}

	if strings.Contains(r.RequestURI, "/invite/forward") {
		r.Header.Set("Content-Type", "application/json")
	}

	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		logger.Info("ctype:", r.Header.Get("Content-Type"))
		return
	}

	body, _ := io.ReadAll(r.Body)
	logger.Info("[%v] %s parameter: %s, %s", id, r.RequestURI,
		(string)(body), http_info(r))
	r.Header.Add("request-parameter", (string)(body))
	r.Body = io.NopCloser(bytes.NewBuffer(body))
}

func RequestOutLog(c *gin.Context) {
	r := c.Request
	if _, ok := apiLogBlackList[r.RequestURI]; ok {
		return
	}

	uri := r.RequestURI
	resp, ok := c.Get("response")

	if !ok || resp == nil || uri == "/" {
		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return
	}
	connId, _ := c.Get("Conn-ID")
	c.Request.Header.Set("Conn-ID", fmt.Sprintf("%v", connId))

	if len(body) > 1024 {
		body = body[:1024]
	}
	logger.InfoCtx(c, "[%v] request: %v; response: %s, %s", connId, uri, string(body), http_info(r))
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestInLog(c)
		defer RequestOutLog(c)
		c.Next()
	}
}

func http_info(r *http.Request) string {
	host := ""
	data := strings.Split(r.Host, ":")
	if len(data) > 0 {
		host = data[0]
	}
	str := fmt.Sprintf("device_id: %s, host: %s, user-agent: %s, version: %s, RemoteAddr: %s, x-user-id: %v, x-sign-id: %v",
		r.Header.Get("device_id"), host, r.Header.Get("user-agent"),
		r.Header.Get("ios-version")+r.Header.Get("android-version"), r.Header.Get("X-Forwarded-For"),
		r.Header.Get("x-user-id"), r.Header.Get("x-sign-id"))
	return str
}
