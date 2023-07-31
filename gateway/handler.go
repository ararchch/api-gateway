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

	service := c.Param("service")
	version := c.Param("version")
	method := c.Param("method")

	kitexClient, err := utils.GenerateClient(service, version)
	if (err != nil) {
		fmt.Println("could not locate service")
		panic(err)
	}

	resp, err := utils.MakeRpcRequest(ctx, kitexClient, method, string(c.Request.Body()))
	if (err != nil){
		fmt.Println("RPC request error")
		panic(err)
	}

	var obj interface{}
	json.Unmarshal([]byte(resp.(string)), &obj)

	c.JSON(consts.StatusOK, obj)
}
