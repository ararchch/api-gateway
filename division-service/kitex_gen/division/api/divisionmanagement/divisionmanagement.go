// Code generated by Kitex v0.6.1. DO NOT EDIT.

package divisionmanagement

import (
	"context"
	api "github.com/ararchch/api-gateway/division-service/kitex_gen/division/api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return divisionManagementServiceInfo
}

var divisionManagementServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "DivisionManagement"
	handlerType := (*api.DivisionManagement)(nil)
	methods := map[string]kitex.MethodInfo{
		"divideNumbers": kitex.NewMethodInfo(divideNumbersHandler, newDivisionManagementDivideNumbersArgs, newDivisionManagementDivideNumbersResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func divideNumbersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.DivisionManagementDivideNumbersArgs)
	realResult := result.(*api.DivisionManagementDivideNumbersResult)
	success, err := handler.(api.DivisionManagement).DivideNumbers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDivisionManagementDivideNumbersArgs() interface{} {
	return api.NewDivisionManagementDivideNumbersArgs()
}

func newDivisionManagementDivideNumbersResult() interface{} {
	return api.NewDivisionManagementDivideNumbersResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) DivideNumbers(ctx context.Context, req *api.DivisionRequest) (r *api.DivisionResponse, err error) {
	var _args api.DivisionManagementDivideNumbersArgs
	_args.Req = req
	var _result api.DivisionManagementDivideNumbersResult
	if err = p.c.Call(ctx, "divideNumbers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
