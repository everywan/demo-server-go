package cmd

import (
	"context"
	"net/http"

	"github.com/everywan/demo-server-go/commons/app"
	"github.com/everywan/demo-server-go/internal/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "http server",
	Long:  `二级命令, 用于启动一个 http server.`,
	Run: func(cmd *cobra.Command, args []string) {
		boot := bootstrap.NewBootstrap()
		defer boot.Teardown()

		logger := boot.GetLogger()
		app := app.New(app.Name("demo_server"),
			app.WithLogger(logger))

		recordCtl := boot.GetRecordController()

		e := gin.Default()
		e.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "enjoy yourself!")
		})

		v1 := e.Group("/v1")
		{
			record := v1.Group("/record")
			record.GET("/:id", recordCtl.Get)
			record.POST("/", recordCtl.Create)
		}
		app.AddBundle()
		app.Run(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
