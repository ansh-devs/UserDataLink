package gapi

import (
	"context"

	"github.com/ansh-devs/tc-assessment/protos"
)

// GetUserByCriteria implements protos.UserServiceServer.
func (b *UserService) GetUserByCriteria(context.Context, *protos.GetUserByCriteriaRequest) (*protos.GetUserByCriteriaResponse, error) {
	panic("unimplemented")
}

// GetUserById implements protos.UserServiceServer.
func (b *UserService) GetUserById(context.Context, *protos.GetUserByIdRequest) (*protos.GetUserByIdResponse, error) {
	panic("unimplemented")
}

// GetUsersListByIds implements protos.UserServiceServer.
func (b *UserService) GetUsersListByIds(context.Context, *protos.GetUsersListByIdsRequest) (*protos.GetUsersListByIdsResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedUserServiceServer implements protos.UserServiceServer.
func (b *UserService) mustEmbedUnimplementedUserServiceServer() {
	panic("unimplemented")
}
