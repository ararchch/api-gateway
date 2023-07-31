// Code generated by Kitex v0.6.1. DO NOT EDIT.

package namemanagement

import (
	"context"
	management "github.com/ararchch/api-gateway/microservices/Name-service/kitex_gen/name/management"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return nameManagementServiceInfo
}

var nameManagementServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "NameManagement"
	handlerType := (*management.NameManagement)(nil)
	methods := map[string]kitex.MethodInfo{
		"helloName": kitex.NewMethodInfo(helloNameHandler, newNameManagementHelloNameArgs, newNameManagementHelloNameResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "management",
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

func helloNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*management.NameManagementHelloNameArgs)
	realResult := result.(*management.NameManagementHelloNameResult)
	success, err := handler.(management.NameManagement).HelloName(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNameManagementHelloNameArgs() interface{} {
	return management.NewNameManagementHelloNameArgs()
}

func newNameManagementHelloNameResult() interface{} {
	return management.NewNameManagementHelloNameResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) HelloName(ctx context.Context, req *management.NameRequest) (r *management.NameResponse, err error) {
	var _args management.NameManagementHelloNameArgs
	_args.Req = req
	var _result management.NameManagementHelloNameResult
	if err = p.c.Call(ctx, "helloName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
