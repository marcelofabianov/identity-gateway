package grpc

import (
	pb "github.com/marcelofabianov/identity-gateway/api/v1/gen"
)

type HealthzServiceServer struct {
	pb.UnimplementedHealthzServer
}

func NewHealthzServiceServer() *HealthzServiceServer {
	return &HealthzServiceServer{}
}
