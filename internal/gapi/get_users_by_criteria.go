package gapi

import (
	"context"

	"github.com/ansh-devs/tc-assessment/internal/validation"
	"github.com/ansh-devs/tc-assessment/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUserByCriteria implements protos.UserServiceServer.
func (s *UserService) GetUsersByCriteria(_ context.Context, req *protos.GetUsersByCriteriaRequest) (*protos.GetUsersByCriteriaResponse, error) {
	err := validation.IsValidCriteria(req.GetType().String())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var usersDTO []*protos.User
	usersDAO, err := s.Repository.GetUsersByField(req.GetType().String(), req.GetValue())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
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
	return &protos.GetUsersByCriteriaResponse{Users: usersDTO}, nil
}
