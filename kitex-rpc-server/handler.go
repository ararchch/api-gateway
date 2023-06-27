package main

import (
	"context"
	"fmt"
	"strconv"

	management "github.com/ararchch/api-gateway/kitex-rpc-server/kitex_gen/addition/management"
)

// AdditionManagementImpl implements the last service interface defined in the IDL.
type AdditionManagementImpl struct{}

// AddNumbers implements the AdditionManagementImpl interface.
func (s *AdditionManagementImpl) AddNumbers(ctx context.Context, req *management.AdditionRequest) (resp *management.AdditionResponse, err error) {

	fmt.Print("HERE----------->")
	fmt.Print(req.FirstNum)
	fmt.Print(req.SecondNum)
	fmt.Print("<------------------")
	
	firstNumInt, err := strconv.Atoi(req.FirstNum)
	if err != nil {
		panic(err)
	}

	secondNumInt, err := strconv.Atoi(req.SecondNum)
	if err != nil {
		panic(err)
	}

	finalSum := firstNumInt + secondNumInt	

	return &management.AdditionResponse{
		Sum: fmt.Sprintf("%d", finalSum),
	}, nil
}
