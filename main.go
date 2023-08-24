package main

import (
	"fmt"
	"grpc-layout/configs"
	"grpc-layout/handler"
	"grpc-layout/helpers/log"
	"grpc-layout/helpers/middleware"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

var (
	gitVersion string = "v0.0.0-master+$Format:%H$"
	gitCommit  string = "$Format:%H$"
	buildDate  string = "1970-01-01T00:00:00Z"
)

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("grpc-layout server, GitVersion: %s, GitCommit: %s, BuildDate: %s\n", gitVersion, gitCommit, buildDate)
	}
	app := cli.App{
		Name:    "grpc-layout",
		Usage:   "grpc-layout http api",
		Version: "v0.0.1",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Value: false,
				Usage: "enable debug mode",
			},
			&cli.StringFlag{
				Name:  "grpc-addr",
				Value: ":8089",
				Usage: "grpc listen addr",
			},
			&cli.StringFlag{
				Name:  "http-addr",
				Value: ":8088",
				Usage: "http listen addr",
			},
			&cli.StringFlag{
				Name:    "parameter",
				Value:   "",
				Usage:   "parameter 其他参数",
				EnvVars: []string{"PARAMETER"}, //全局变量的设置方式
			},
		},
		Before: func(ctx *cli.Context) error {
			if !ctx.Bool("debug") {
				log.SetOutput(os.Stderr, log.LevelDebug, false)
				// log.SetOutput(&log.LogDir{
				// 	Dir:    "./log_doc",
				// 	Format: "20060102_15",
				// }, log.LevelInfo, false)
			} else {
				log.SetOutput(os.Stderr, log.LevelDebug, true)
			}
			if err := configs.Load(ctx); err != nil {
				log.Info("load args err")
				return err
			}
			return nil
		},
		Action: func(ctx *cli.Context) error {
			log.Debug("serve start ...")
			grpc_addr := ctx.String("grpc-addr")
			http_addr := ctx.String("http-addr")
			lis, err := net.Listen("tcp", grpc_addr)
			if err != nil {
				return fmt.Errorf("failed to listen: %v", err)
			}
			//服务实例
			api, err := handler.NewApp()
			if err != nil {
				return fmt.Errorf("new App err: %v", err)
			}
			//grpc server
			grpc_server := grpc.NewServer(grpc.ChainUnaryInterceptor(middleware.LogUnaryServerInterceptor()))
			api.RegisterAppServer(grpc_server)
			reflection.Register(grpc_server)
			go func() {
				log.Info(fmt.Sprintf("listen grpc %v", lis.Addr()))
				if err := grpc_server.Serve(lis); err != nil {
					log.Fatal("start grpc server err")
				}
			}()
			// grpc gw
			mux := runtime.NewServeMux()
			err = api.RegisterAppFromEndpoint(mux, grpc_addr)
			if err != nil {
				return err
			}
			log.Info(fmt.Sprintf("listen http %v", http_addr))
			// Start HTTP server (and proxy calls to gRPC server endpoint)
			if err := http.ListenAndServe(http_addr, mux); err != nil {
				return err
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Info(err.Error())
	}
}
