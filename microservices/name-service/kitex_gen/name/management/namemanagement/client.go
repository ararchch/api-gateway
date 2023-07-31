// Code generated by Kitex v0.6.1. DO NOT EDIT.

package namemanagement

import (
	"context"
	management "github.com/ararchch/api-gateway/microservices/name-service/kitex_gen/name/management"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	HelloName(ctx context.Context, req *management.NameRequest, callOptions ...callopt.Option) (r *management.NameResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kNameManagementClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kNameManagementClient struct {
	*kClient
}

func (p *kNameManagementClient) HelloName(ctx context.Context, req *management.NameRequest, callOptions ...callopt.Option) (r *management.NameResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HelloName(ctx, req)
}
