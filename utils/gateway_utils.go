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

// generates and returns a new kitex client based on input parameters
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

	// creates file name using the version following pre-defined convention
	idlFileName := fmt.Sprintf("%s.thrift", version)

	// gets IDL from the [service name] branch of the IDL repo
	idlUrl := fmt.Sprintf("https://raw.githubusercontent.com/ararchch/api-gateway-idl/%s/%s", serviceName, idlFileName)

	// accesses latest version of IDL file from github repo
	resp, err := http.Get(idlUrl)
	if err != nil {
		fmt.Println("Error in accessing IDL file")
	}
	defer resp.Body.Close()

	// extracts bytes from response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("Error in Reading IDL file\n")
		panic(err)
	}

	// name of new temp file to be created
	tempIdlName := fmt.Sprintf("%s-%s", serviceName, idlFileName)
	// accesses temp directory of OS of system the server is running on
	tempDir := os.TempDir()
	// creates full local IDL path
	idlPath := filepath.Join(tempDir, tempIdlName)
	// writes temp IDL file with same data extracted from the github IDL file
	err = os.WriteFile(idlPath, body, 0644)
	if err != nil {
		fmt.Println("Error in creating temporary IDL file")
		panic(err)
	}

	// importing idl for reference(generic call)
	p, err := generic.NewThriftFileProvider(idlPath)
	if err != nil {
		panic(err)
	}

	// convert to thrift generic form
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	//creates kitex client options array
	var options []kitexClient.Option

	// adds service discovery, load balancing capabilities to options array (these 'options' are permanently integrated)
	options = append(options,
		kitexClient.WithResolver(r),
		kitexClient.WithLoadBalancer(lb),
	)

	// appends any passed in options from request side to the gateway
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

	// returns client to caller
	return client, nil
}

// converts item to json string format
func jsonStringify(item any) (string, error) {
	jsonForm, err := json.Marshal(&item)
	if err != nil {
		panic(err)
	}

	return string(jsonForm), nil
}

// makes rpc request to kitex client based on input parameters
func MakeRpcRequest(ctx context.Context, kitexClient genericclient.Client, methodName string, request string) (interface{}, error) {

	// making generic call to specified method of client
	respRpc, err := kitexClient.GenericCall(ctx, methodName, request)
	if err != nil {
		panic(err)
	}
	
	// return to client
	return respRpc, nil
}

func ReadIdlFromGithub(service string, version string) (string, error) {
	// creates file name using the version following pre-defined convention
	idlFileName := fmt.Sprintf("%s.thrift", version)

	// gets IDL from the [service name] branch of the IDL repo
	idlUrl := fmt.Sprintf("https://raw.githubusercontent.com/ararchch/api-gateway-idl/%s/%s", service, idlFileName)

	// Download the IDL file
	resp, err := http.Get(idlUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// extracts bytes from response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("Error in Reading IDL file\n")
		panic(err)
	}

	// name of new temp file to be created
	tempIdlName := fmt.Sprintf("%s-%s", service, idlFileName)
	// accesses temp directory of OS of system the server is running on
	tempDir := os.TempDir()
	// creates full local IDL path
	idlPath := filepath.Join(tempDir, tempIdlName)
	// writes temp IDL file with same data extracted from the github IDL file
	err = os.WriteFile(idlPath, body, 0644)
	if err != nil {
		fmt.Println("Error in creating temporary IDL file")
		panic(err)
	}

	return idlPath, nil
}