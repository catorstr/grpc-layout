package middleware

//认证
import (
	"context"
	"grpc-layout/helpers/middleware/jwt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthToken(ctx context.Context) error {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		app_id  string
		app_key string
	)
	if val, ok := md["app_id"]; ok {
		app_id = val[0]
	}
	if val, ok := md["app_key"]; ok {
		app_key = val[0]
	}
	mc, err := jwt.ParseToken(app_key, []byte("secretKey"))
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Token认证信息无效: app_id=%s, app_key=%s", app_id, app_key)
	}
	metadata.AppendToOutgoingContext(ctx, "user_info", mc.Account)
	return nil
}
