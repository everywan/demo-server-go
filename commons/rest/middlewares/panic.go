package middlewares

import (
	"github.com/everywan/demo-server-go/commons/errors"
	"github.com/everywan/demo-server-go/commons/rest"
	"github.com/everywan/demo-server-go/commons/utils"
	"github.com/gin-gonic/gin"
	pkgErrors "github.com/pkg/errors"
)

func PanicAsError(g *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.SafeRun(g.Next)
		if err != nil {
			switch realErr := pkgErrors.Cause(err).(type) {
			case *errors.ErrorCode:
				rest.FailResponse()
				// rest.RenderError(w, realErr)
				return
			default:
				// panic(realErr)
			}
		}
	}
}
