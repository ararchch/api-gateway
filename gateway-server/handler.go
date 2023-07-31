package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/ararchch/api-gateway/utils"
)

// Handles incoming requests to server
func HandlePostRequest(ctx context.Context, c *app.RequestContext) {

	// parsing HTTP request path for relevant data
	service := c.Param("service") // RPC service name (same as the one used to register service / create IDL branch)
	version := c.Param("version") // IDL version (same name as IDL file stored in IDL repo under service branch)
	method := c.Param("method")	// method to be called within service

	// generating a kitex client using the utils package's GenerateClient method
	kitexClient, err := utils.GenerateClient(service, version, utils.ConnectionTimeout(1000), utils.RpcTimeout(1000))
	if (err != nil) {
		fmt.Println("could not locate service")
		panic(err)
	}

	// calls utils Make RpcRequest method 
	resp, err := utils.MakeRpcRequestWithRetry(ctx, kitexClient, method, string(c.Request.Body()), 3)
	if (err != nil){
		fmt.Println("RPC request error")
		panic(err)
	}

	// object to be returned to HTTP request sender
	var httpReturnObj interface{}

	// unmarshall json string to returned obj
	json.Unmarshal([]byte(resp.(string)), &httpReturnObj)

	// return to client as JSON response
	c.JSON(consts.StatusOK, httpReturnObj)
}
