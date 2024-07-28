package grpc

import (
	"context"

	pb "github.com/marcelofabianov/identity-gateway/api/v1/gen"
)

type HealthzServiceServer struct {
	pb.UnimplementedHealthzServer
}

func NewHealthzServiceServer() *HealthzServiceServer {
	return &HealthzServiceServer{}
}

func (s *HealthzServiceServer) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	status := "OK"
	message := "Service is healthy"

	return &pb.CheckResponse{
		Status:  status,
		Message: message,
	}, nil
}
