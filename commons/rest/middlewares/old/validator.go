package middleware

import (
	"github.com/gin-gonic/gin"
)

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Next()
	}
}
