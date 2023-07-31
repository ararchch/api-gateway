package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	kitexClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
)

// adds RPC timeout functionality to gateway, returns kitexClient.option type that can be passed as optional input to 'GenerateClient' method
func RpcTimeout(dur int) kitexClient.Option {
	return kitexClient.WithRPCTimeout(time.Duration(dur) * time.Millisecond)
}

// adds Connection timeout functionality to gateway, returns kitexClient.option type that can be passed as optional input to 'GenerateClient' method
func ConnectionTimeout(dur int) kitexClient.Option {
	return kitexClient.WithConnectTimeout(time.Duration(dur) * time.Millisecond)
}

// same as MakeRpcRequest method, except it retries upon failure a provided number of times
func MakeRpcRequestWithRetry(ctx context.Context, kitexClient genericclient.Client, methodName string, request string, retryCount int) (interface{}, error) {
	
	var errResp error
	// looping up to 1 + retryCount times
	for i := 0; i < retryCount + 1; i++ {

		// prints out retry count
		if i > 0 {
			fmt.Printf("------> Retrying.... %d \n", i)
		}
		
		// Making generic call to the specified method of the client
		respRpc, err := kitexClient.GenericCall(ctx, methodName, request)
		if err != nil {
			// Retry on error
			errResp = err
			continue
		}

		// Return response, nil if the request succeeded
		return respRpc, nil

	}

	// Returns nil, the last error if all retries failed
	if errResp != nil {
		fmt.Printf("Failed despite %d retries \n", retryCount)
		return nil, errResp
	}

	// If no retries were done, return nil, a generic error
	return nil, errors.New("retry failed")
}
