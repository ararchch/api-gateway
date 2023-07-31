// Code generated by hertz generator.

package handler

import (
	"context"
	"fmt"
	//"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	//"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	gatewayUtils "github.com/ararchch/api-gateway/utils"
)

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {

	service := c.Param("service")
	method := c.Param("method")

	//fmt.Println(string(c.Request.Body()))

	kitexClient, err := gatewayUtils.GenerateClient(service)
	if (err != nil) {
		fmt.Println("could not locate service\n")
		panic(err)
	}

	// type AdditionRequest struct {
	// 	FirstNum  string 
	// 	SecondNum string 
	// }
	// var req AdditionRequest

	// type AdditionResponse struct {
	// 	Sum string 
	// }
	//var resp AdditionResponse

	// bind error params to req (pre-generated)
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	c.String(consts.StatusBadRequest, err.Error())
	// 	return
	// }

	resp, err := gatewayUtils.MakeRpcRequest2(ctx, kitexClient, method, string(c.Request.Body()))
	if (err != nil){
		fmt.Println("RPC request error \n")
		panic(err)
	}

	c.JSON(consts.StatusOK, resp)
}
