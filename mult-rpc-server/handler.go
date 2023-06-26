package main

import (
	"context"
	management "github.com/ararchch/api-gateway/mult-rpc-server/kitex_gen/multiplication/management"
)

// MultiplicationManagementImpl implements the last service interface defined in the IDL.
type MultiplicationManagementImpl struct{}

// MultiplyNumbers implements the MultiplicationManagementImpl interface.
func (s *MultiplicationManagementImpl) MultiplyNumbers(ctx context.Context, req *management.MultiplicationRequest) (resp *management.MultiplicationResponse, err error) {

	return &management.MultiplicationResponse{
		Product: req.FirstNum * req.SecondNum,
	}, nil
}
