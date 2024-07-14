package grpc

import (
	"context"
	"time"

	userpb "github.com/marcelofabianov/pejota/internal/user/adapter/grpc/generated"
	"github.com/marcelofabianov/pejota/internal/user/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	createdAt, _ := time.Parse(time.RFC3339, user.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, user.UpdatedAt)

	return &userpb.GetUserResponse{
		PublicId:     user.PublicID,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		LoginEnabled: user.LoginEnabled,
		CreatedAt:    timestamppb.New(createdAt),
		UpdatedAt:    timestamppb.New(updatedAt),
	}, nil
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	input := port.CreateUserInputServiceApplication{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole(),
	}

	user, err := s.application.CreateUser(input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating user: %v", err)
	}

	createdAt, _ := time.Parse(time.RFC3339, user.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, user.UpdatedAt)

	return &userpb.CreateUserResponse{
		PublicId:     user.PublicID,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		LoginEnabled: user.LoginEnabled,
		CreatedAt:    timestamppb.New(createdAt),
		UpdatedAt:    timestamppb.New(updatedAt),
	}, nil
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	publicID := req.GetPublicId()

	input := port.DeleteUserInputServiceApplication{
		PublicID: publicID,
	}

	user, err := s.application.DeleteUser(input)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found with public_id: %s", publicID)
	}

	return &userpb.DeleteUserResponse{
		Success: user.Success,
	}, nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	publicID := req.GetPublicId()

	input := port.UpdateUserInputServiceApplication{
		PublicID: publicID,
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Role:     req.GetRole(),
	}

	user, err := s.application.UpdateUser(input)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found with public_id: %s", publicID)
	}

	createdAt, _ := time.Parse(time.RFC3339, user.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, user.UpdatedAt)

	return &userpb.UpdateUserResponse{
		PublicId:     user.PublicID,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		LoginEnabled: user.LoginEnabled,
		CreatedAt:    timestamppb.New(createdAt),
		UpdatedAt:    timestamppb.New(updatedAt),
	}, nil
}
