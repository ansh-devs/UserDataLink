package gapi

import (
	"context"

	"github.com/ansh-devs/tc-assessment/internal/validation"
	"github.com/ansh-devs/tc-assessment/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUserById implements protos.UserServiceServer.
func (s *UserService) GetUserById(_ context.Context, req *protos.GetUserByIdRequest) (*protos.GetUserByIdResponse, error) {
	if err := validation.IsValidId(req.GetId()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user, err := s.Repository.GetUserById(int64(req.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protos.GetUserByIdResponse{
		User: &protos.User{
			Id:     int64(user.ID),
			FName:  user.FName,
			Height: float32(user.Height),
			City:   user.City,
			Phone:  user.Phone,
		}}, nil
}
