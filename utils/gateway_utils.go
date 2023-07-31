package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	kitexClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func GenerateClient(serviceName string, version string, opts ...kitexClient.Option) (genericclient.Client, error){

	// inital declarations
	var err error

	// initating loadbalancer
	lb := loadbalance.NewWeightedBalancer()

	// initating etcs resolver (for service discovery)
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	idlUrl := fmt.Sprintf("https://raw.githubusercontent.com/ararchch/api-gateway-idl/%s/%s", serviceName, version)

	resp, err := http.Get(idlUrl)
	if err != nil {
		fmt.Println("Error in accessing IDL file\n")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("Error in Reading IDL file\n")
		panic(err)
	}

	tempIdlName := fmt.Sprintf("%s-%s.thrift", serviceName, version)
	tempDir := os.TempDir()
	filename := filepath.Join(tempDir, tempIdlName)
	err = os.WriteFile(filename, body, 0644)
	if err != nil {
		fmt.Println("Error in creating temporary IDL file")
		panic(err)
	}

	// importing idl for reference(generic call)
	p, err := generic.NewThriftFileProvider(filename)
	if err != nil {
		panic(err)
	}

	// convert to thrift generic form
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	var options []kitexClient.Option
	options = append(options,
		kitexClient.WithResolver(r),
		kitexClient.WithLoadBalancer(lb),
	)
	options = append(options, opts...)

	// create new generic client
	client, err := genericclient.NewClient(
		serviceName,
		g,
		options...
	)
	if err != nil {
		panic(err)
	}

	return client, nil
}

func jsonStringify(item any) (string, error) {
	// convert to request struct to JSON format (so it can be converted to json string)
	jsonForm, err := json.Marshal(&item)
	if err != nil {
		panic(err)
	}

	return string(jsonForm), nil
}

func MakeRpcRequest(ctx context.Context, kitexClient genericclient.Client, methodName string, request string) (interface{}, error) {

	// making generic call to addNumbers method of client
	respRpc, err := kitexClient.GenericCall(ctx, methodName, request)
	if err != nil {
		panic(err)
	}

	fmt.Println(respRpc)

	return respRpc, nil
}