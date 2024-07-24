package rpc

import (
	"context"
	"hex_arch_go/internal/adapters/framework/left/grpc/proto/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetNumber1() == 0 || req.GetNumber2() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetAddition(req.Number1, req.Number2)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{Result: answer}

	return ans, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetNumber1() == 0 || req.GetNumber2() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetSubtraction(req.Number1, req.Number2)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{Result: answer}

	return ans, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetNumber1() == 0 || req.GetNumber2() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetMultiplication(req.Number1, req.Number2)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{Result: answer}

	return ans, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetNumber1() == 0 || req.GetNumber2() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetDivision(req.Number1, req.Number2)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{Result: answer}

	return ans, nil
}
