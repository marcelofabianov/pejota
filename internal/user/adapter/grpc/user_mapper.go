package grpc

import (
	"time"

	userpb "github.com/marcelofabianov/pejota/internal/user/adapter/grpc/generated"
	"github.com/marcelofabianov/pejota/internal/user/port"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapToUserProto(user port.GetUserOutputServiceApplication) *userpb.GetUserResponse {
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
	}
}
