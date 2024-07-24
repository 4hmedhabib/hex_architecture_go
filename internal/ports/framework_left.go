package ports

import (
	"context"
	"hex_arch_go/internal/adapters/framework/left/grpc/proto/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
}
