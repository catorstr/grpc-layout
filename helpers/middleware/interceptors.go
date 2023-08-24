package middleware

// interceptor 拦截器
import (
	"context"
	"grpc-layout/helpers/log"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var (
			grpcgateway_user_agent   string
			grpcgateway_content_type string
			x_forwarded_for          string
		)
		md, _ := metadata.FromIncomingContext(ctx)
		if val, ok := md["grpcgateway-user-agent"]; ok {
			grpcgateway_user_agent = val[0]
		}
		if val, ok := md["grpcgateway-content-type"]; ok {
			grpcgateway_content_type = val[0]
		}

		if val, ok := md["x-forwarded-for"]; ok {
			x_forwarded_for = val[0]
		}
		log.Info(
			"LogUnaryServerInterceptor",
			slog.String("grpcgateway-content-type", grpcgateway_content_type),
			slog.String("grpcgateway-user-agent", grpcgateway_user_agent),
			slog.String("x-forwarded-for", x_forwarded_for),
			slog.String("method", info.FullMethod),
		)
		// 继续处理请求
		return handler(ctx, req)
	}
}

// func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	err := auth(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 继续处理请求
// 	return handler(ctx, req)
// }
