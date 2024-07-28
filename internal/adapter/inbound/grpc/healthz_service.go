package grpc

import (
	"context"

	pb "github.com/marcelofabianov/identity-gateway/api/v1/gen"
)

type HealthzServiceServer struct {
	pb.UnimplementedHealthzServiceServer
}

func NewHealthzServiceServer() *HealthzServiceServer {
	return &HealthzServiceServer{}
}

func (s *HealthzServiceServer) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}
