package gapi

import (
	"context"

	"github.com/ansh-devs/tc-assessment/protos"
)

// GetUsersListByIds implements protos.UserServiceServer.
func (b *UserService) GetUsersListByIds(_ context.Context, req *protos.GetUsersListByIdsRequest) (resp *protos.GetUsersListByIdsResponse, err error) {
	panic("unimplemented")
}
