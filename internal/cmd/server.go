package cmd

import (
	"net/http"

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

		// logger := boot.GetLogger()
		recordCtl := boot.GetRecordController()

		e := gin.Default()
		// e.Use(middleware.Logger())
		// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// 	AllowOrigins:     opts.Server.CorsAllowOrigin,
		// 	AllowCredentials: true,
		// }))

		e.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "enjoy yourself!")
		})

		v1 := e.Group("/v1")
		{
			record := v1.Group("/record")
			record.GET("/:id", recordCtl.Get)
			record.Post("/", recordCtl.Create)
		}

		// quit := make(chan os.Signal, 1)
		// go func() {
		// 	// 当程序较多/HTTP设置较多时, 可以单独封装Server组件, 在组件内计算这些值
		// 	address := fmt.Sprintf("%s:%d", opts.Server.Host, opts.Server.Port)
		// 	err = e.Start(address)
		// 	if err != nil {
		// 		logger.Fatal("start echo error, error is ", err)
		// 		quit <- os.Interrupt
		// 	}
		// }()
		// signal.Notify(quit, os.Interrupt)
		// <-quit

		// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancel()
		// if err := e.Shutdown(ctx); err != nil {
		// 	e.Logger.Fatal(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
