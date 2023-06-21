package main

import (
	"context"
	management "github.com/ararchch/api-gateway/kitex-rpc-server/kitex_gen/addition/management"
)

// AdditionManagementImpl implements the last service interface defined in the IDL.
type AdditionManagementImpl struct{}

// AddNumbers implements the AdditionManagementImpl interface.
func (s *AdditionManagementImpl) AddNumbers(ctx context.Context, req *management.AdditionRequest) (resp *management.AdditionResponse, err error) {
	
	return &management.AdditionResponse{
		Sum: req.FirstNum + req.SecondNum,
	}
}
