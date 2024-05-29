package gapi

import (
	"context"

	"github.com/ansh-devs/tc-assessment/protos"
)

// GetUserById implements protos.UserServiceServer.
func (b *UserService) GetUserById(_ context.Context, req *protos.GetUserByIdRequest) (resp *protos.GetUserByIdResponse, err error) {
	panic("unimplemented")
}
