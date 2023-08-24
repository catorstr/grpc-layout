package handler

import (
	"context"
	"grpc-layout/api/greeter"
	"grpc-layout/handler/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	//服务聚合
	Greeter *service.GreeterService
}

func NewApp() (*App, error) {
	return &App{
		Greeter: &service.GreeterService{},
	}, nil
}

// 注册grpc服务
func (f *App) RegisterAppServer(conn *grpc.Server) {
	greeter.RegisterGreeterServer(conn, f.Greeter)
}

// 使grpc代理http
func (f *App) RegisterAppFromEndpoint(mux *runtime.ServeMux, grpc_addr string) (err error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx := context.Background()
	// err = pb.RegisterYourServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	greeter.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpc_addr, opts)
	return nil
}
