package router

import (
	"net/http"

	"github.com/everywan/demo-server-go/internal/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewRouter(boot *bootstrap.Bootstrap) *gin.Engine {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "enjoy yourself!")
	})

	v1 := e.Group("/v1")
	{
		recordCtl := boot.GetRecordController()
		record := v1.Group("/record")
		record.GET("/:id", recordCtl.Query)
	}
	return e
}
