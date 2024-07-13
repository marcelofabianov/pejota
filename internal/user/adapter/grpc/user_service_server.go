package grpc

import (
	"context"

	userpb "github.com/marcelofabianov/pejota/internal/user/adapter/grpc/generated"
	"github.com/marcelofabianov/pejota/internal/user/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	application port.UserServiceApplication
}

func NewUserServiceServer(application port.UserServiceApplication) *UserServiceServer {
	return &UserServiceServer{
		application: application,
	}
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	publicID := req.GetPublicId()

	input := port.GetUserInputServiceApplication{
		PublicID: publicID,
	}

	user, err := s.application.GetUser(input)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found with public_id: %s", publicID)
	}

	userProto := mapToUserProto(user)

	return userProto, nil
}
