package service

import (
	"context"

	pb "grpc-layout/api/greeter"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

func (s *GreeterService) PostSayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		QaUUid: "post say hello",
	}, nil
}
func (s *GreeterService) GetSayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		QaUUid: "get say hello",
	}, nil
}
