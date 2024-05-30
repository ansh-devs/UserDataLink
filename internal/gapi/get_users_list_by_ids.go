package gapi

import (
	"context"

	"github.com/ansh-devs/tc-assessment/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUsersListByIds implements protos.UserServiceServer.
func (s *UserService) GetUsersListByIds(_ context.Context, req *protos.GetUsersListByIdsRequest) (*protos.GetUsersListByIdsResponse, error) {
	repoUsers, err := s.Repository.GetUsersListByIds(req.Ids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var users []*protos.User
	for _, v := range repoUsers {
		users = append(users, &protos.User{
			Id:      v.ID,
			FName:   v.FName,
			Height:  float32(v.Height),
			City:    v.City,
			Phone:   v.Phone,
			Married: v.Married,
		})
	}
	return &protos.GetUsersListByIdsResponse{Users: users}, nil
}
