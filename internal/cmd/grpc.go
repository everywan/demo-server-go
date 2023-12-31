package cmd

import (
	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "grpc server",
	Long:  `二级命令, 用于启动一个 grpc server.`,
	Run: func(cmd *cobra.Command, args []string) {
		// opts, err := loadOptions()
		// handleInitError("load_options", err)
		// boot := newBootstrap(opts)
		// defer boot.Teardown()

		// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.Grpc.Port))
		// // handleInitError("grpc_listen", err)

		// demoCtl := controllers.NewDemoController(boot.GetLogger(), boot.GetDemoSvc())

		// logger := boot.GetLogger().WithField("scope", "demo")
		// gs := grpc.NewServer(
		// 	grpc.KeepaliveParams(keepalive.ServerParameters{
		// 		Time: 5 * time.Second,
		// 	}),
		// 	grpc_middleware.WithUnaryServerChain(
		// 		grpc_ctxtags.UnaryServerInterceptor(
		// 			grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
		// 			grpc_ctxtags.WithFieldExtractor(func(fullMethod string, req interface{}) map[string]interface{} {
		// 				fields := map[string]interface{}{"request_id": xid.New().String()}
		// 				return fields
		// 			}),
		// 		),
		// 		grpc_logrus.UnaryServerInterceptor(logger),
		// 		grpc_logrus.PayloadUnaryServerInterceptor(logger, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true }),
		// 		grpc_validator.UnaryServerInterceptor(),
		// 	),
		// )
		// // 不同的service分别依次注册
		// pb.RegisterDemoServiceServer(gs, demoCtl)

		// quit := make(chan os.Signal, 1)
		// go func() {
		// 	logger.Infof("grpc server start at port %d...", opts.Grpc.Port)
		// 	err = gs.Serve(lis)
		// 	if err != nil {
		// 		logger.Fatalf("start server error, error is %v ", err)
		// 		quit <- os.Interrupt
		// 	}
		// }()
		// signal.Notify(quit, os.Interrupt)
		// <-quit

		// gs.GracefulStop()
	},
}

type GrpcOption struct {
	Port int `mapstructure:"port" yaml:"port"`
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
