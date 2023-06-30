package main

import (
	"context"
	"fmt"
	"strconv"

	management "github.com/ararchch/api-gateway/microservices/multiplication-service/kitex_gen/multiplication/management"
)

// MultiplicationManagementImpl implements the last service interface defined in the IDL.
type MultiplicationManagementImpl struct{}

// MultiplyNumbers implements the MultiplicationManagementImpl interface.
func (s *MultiplicationManagementImpl) MultiplyNumbers(ctx context.Context, req *management.MultiplicationRequest) (resp *management.MultiplicationResponse, err error) {

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

	// multiply two numbers together
	finalProduct := firstNumInt * secondNumInt

	// convert finalSum to string and return response of type MultiplicationResponse and error = nil
	return &management.MultiplicationResponse{
		Product: fmt.Sprintf("%d", finalProduct),
	}, nil

}
