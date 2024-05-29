package gapi

import (
	"context"

	"github.com/ansh-devs/tc-assessment/protos"
)

// GetUserByCriteria implements protos.UserServiceServer.
func (b *UserService) GetUserByCriteria(_ context.Context, req *protos.GetUsersByCriteriaRequest) (resp *protos.GetUsersByCriteriaResponse, err error) {
	panic("unimplemented")
}
