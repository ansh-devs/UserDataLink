package gapi

import (
	"context"

	"github.com/ansh-devs/tc-assessment/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUsersListByIds implements protos.UserServiceServer.
func (s *UserService) GetUsersListByIds(_ context.Context, req *protos.GetUsersListByIdsRequest) (*protos.GetUsersListByIdsResponse, error) {
	usersDAO, err := s.Repository.GetUsersListByIds(req.Ids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var usersDTO []*protos.User
	for _, userDAO := range usersDAO {
		usersDTO = append(usersDTO, &protos.User{
			Id:      userDAO.ID,
			FName:   userDAO.FName,
			Height:  float32(userDAO.Height),
			City:    userDAO.City,
			Phone:   userDAO.Phone,
			Married: userDAO.Married,
		})
	}
	return &protos.GetUsersListByIdsResponse{Users: usersDTO}, nil
}
