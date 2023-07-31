package main

import (
	"context"
	management "github.com/ararchch/api-gateway/microservices/Name-service/kitex_gen/name/management"
)

// NameManagementImpl implements the last service interface defined in the IDL.
type NameManagementImpl struct{}

// HelloName implements the NameManagementImpl interface.
func (s *NameManagementImpl) HelloName(ctx context.Context, req *management.NameRequest) (resp *management.NameResponse, err error) {
	// TODO: Your code here...
	return &management.NameResponse{
		HelloResp: "Hello, " + req.Name,
	}, nil 

}
