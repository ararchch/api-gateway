// Code generated by hertz generator.

package api

import (
	"context"
	"fmt"

	api "github.com/ararchch/api-gateway/hz-http-server/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	
	kitexClient "github.com/cloudwego/kitex/client"
	"github.com/ararchch/api-gateway/kitex-rpc-server/kitex_gen/addition/management"
	"github.com/ararchch/api-gateway/kitex-rpc-server/kitex_gen/addition/management/additionmanagement"	
)

// AddNumbers .
// @router /add [POST]
func AddNumbers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.AdditionRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	client, err := additionmanagement.NewClient("sum", kitexClient.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	reqRpc := &management.AdditionRequest{
		FirstNum: req.FirstNum,
		SecondNum: req.SecondNum,
	}

	respRpc, err := client.AddNumbers(ctx, reqRpc)
	if err != nil {
		
		panic(err)
	}

	resp := api.AdditionResponse{
		Sum: fmt.Sprintf("%d", respRpc.Sum),
	}

	c.JSON(consts.StatusOK, resp)
}
