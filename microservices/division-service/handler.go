package main

import (
	"context"
	"fmt"
	"strconv"

	api "github.com/ararchch/api-gateway/microservices/division-service/kitex_gen/division/api"
)

// DivisionManagementImpl implements the last service interface defined in the IDL.
type DivisionManagementImpl struct{}

// DivideNumbers implements the DivisionManagementImpl interface.
func (s *DivisionManagementImpl) DivideNumbers(ctx context.Context, req *api.DivisionRequest) (resp *api.DivisionResponse, err error) {
	
	// parse int from string of First Number
	firstNumInt, err := strconv.Atoi(req.FirstNum)
	if err != nil {
		panic(err)
	}

	// parse int from string of Second Number
	secondNumInt, err := strconv.Atoi(req.SecondNum)
	if err != nil {
		panic(err)
	}

	// divide two numbers 
	finalQuotient := firstNumInt / secondNumInt;

	// convert finalSum to string and return response of type DivisionResponse and error = nil
	return &api.DivisionResponse{
		Quotient: fmt.Sprintf("%d", finalQuotient),
	}, nil

}
