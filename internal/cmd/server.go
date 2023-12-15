package cmd

import (
	"context"

	"github.com/everywan/demo-server-go/commons/app"
	"github.com/everywan/demo-server-go/internal/bootstrap"
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

		app.AddBundle()
		app.Run(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
