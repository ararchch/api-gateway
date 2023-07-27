package utils

import (
	"context"
	"encoding/json"
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

func MakeRpcRequestWithRetry(ctx context.Context, kitexClient genericclient.Client, methodName string, request interface{}, response interface{}, retryCount int) error {
	stringedReq, err := jsonStringify(request)
	if err != nil {
		return err
	}

	var errResp error
	for i := 0; i < retryCount; i++ {
		// Making generic call to the specified method of the client
		fmt.Printf("------> Retrying.... %d \n", i)
		respRpc, err := kitexClient.GenericCall(ctx, methodName, stringedReq)
		if err != nil {
			// Retry on error
			errResp = err
			continue
		}

		// Unmarshal the response
		err = json.Unmarshal([]byte(respRpc.(string)), response)
		if err != nil {
			return err
		}

		// Return nil if the request succeeded
		return nil
	}

	// Return the last error if all retries failed
	if errResp != nil {
		fmt.Printf("Failed despite %d retries \n", retryCount)
		return errResp
	}

	// If no retries were done, return a generic error
	return errors.New("retry failed")
}
