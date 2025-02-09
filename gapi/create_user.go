package gapi

import (
	"context"

	"github.com/jpmoraess/appointment-api/internal/usecases"
	"github.com/jpmoraess/appointment-api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	output, err := server.register.Execute(ctx, &usecases.RegisterInput{
		Name:     req.Name,
		Password: req.Password,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register user: %s", err)
	}

	return &pb.CreateUserResponse{
		UserId:   output.UserID,
		TenantId: output.TenantID,
	}, nil
}
