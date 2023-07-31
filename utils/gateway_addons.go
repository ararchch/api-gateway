package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	kitexClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
)

func RpcTimeout(dur int) kitexClient.Option {
	return kitexClient.WithRPCTimeout(time.Duration(dur) * time.Millisecond)
}

func ConnectionTimeout(dur int) kitexClient.Option {
	return kitexClient.WithConnectTimeout(time.Duration(dur) * time.Millisecond)
}

func MakeRpcRequestWithRetry(ctx context.Context, kitexClient genericclient.Client, methodName string, request string, retryCount int) (interface{}, error) {
	
	var errResp error
	for i := 0; i < retryCount + 1; i++ {
		// Making generic call to the specified method of the client
		if i > 0 {
			fmt.Printf("------> Retrying.... %d \n", i)
		}
		
		respRpc, err := kitexClient.GenericCall(ctx, methodName, request)
		if err != nil {
			// Retry on error
			errResp = err
			continue
		}

		fmt.Println(respRpc)

		// Return response, nil if the request succeeded
		return respRpc, nil

	}

	// Return the last error if all retries failed
	if errResp != nil {
		fmt.Printf("Failed despite %d retries \n", retryCount)
		return nil, errResp
	}

	// If no retries were done, return a generic error
	return nil, errors.New("retry failed")
}
