# CloudWeGo API Gateway

An API Gateway written in `GO` that uses CloudWeGo `Kitex` and `Hertz` frameworks to develop an API Gateway that accepts `HTTP` requests encoded in `JSON` format, Utilising the Generic-Call feature of `Kitex` to translate these requests into `Thrift` binary format requests. The API Gateway will then proceed to forward the translated requests to one of the backend `RPC` servers obtained from the registry centre. 

## Tech Design

[Tech Design Link](https://drive.google.com/file/d/1guiZxHbMgkmHLFn2DxIcptSoMvI5CdSa/view?usp=drive_link) 

## API Gateway Design Diagram
![API Gateway Diagram](API-Diagram.png)

## How it works

1. Client sends HTTP request with a JSON body with a "FirstNum" and "SecondNum" attribute. <br>
2. Hertz server processes the request and sends it to the Kitex Servers. <br>
3. A generic Kitex server will be generated and registered in the `etcd` registry. <br>
4. Kitex server will process the request and convert it into `Thrift Binary Format` before sending it to the RPC servers. <br>
5. The RPC servers will process the `addition` request and returns the response. <br>
6. The API gateway will process and encode the response as an HTTP Response.

## Installation

Before running the project, make sure you install the following, <br>
- [Hertz](https://www.cloudwego.io/docs/hertz/getting-started/)<br>
- [Kitex](https://www.cloudwego.io/docs/kitex/getting-started/)<br>
- [etcd](https://etcd.io/docs/v3.2/install/)<br>
- [PostMan](https://www.postman.com/downloads/)

## Instructions to run the API Gateway

Make sure `etcd` is running.

- Instructions [here](https://etcd.io/docs/v3.5/quickstart/) for `Mac` users. If your machine is running on `Windows`, unzip the folder and run the `etcd` application in the folder.

In the `hertz-http-server` directory, run `go run .` in the terminal.

In the `addition-service` directory, run `go run .` in the terminal.

Test with `Postman` using the following command: `http://127.0.0.1:8080/add` with the JSON body of:

```json
{
  "firstnum" : "1",
  "secondnum" : "2"
}
```

Anticipate a response of:

```json
{
  "Sum": "3"
}
```

## Creating a new microservice and integrating it on the API Gateway

- The `hertz-http-server` folder contains the code for accepting `HTTP` requests and the primary business code for the API gateway implementation.

- The `utils` folder contains the file `utils.go`, which offers a number of useful functions to assist developers with using our API gateway.

- The `thrift-idl` folder contains thrift Interface Definition Language files. These files are used to generate the infrastructure code for the `Hertz` `HTTP` server as well as any microservices that are developed.

- The `addition-service` and `multiplication-service` folders contain the code for addition and multiplication RPC servers i.e. addition and multiplication microservices.<br>
These are not part of the gateway itself and are more like examples of how you can use it. Now we will be implementing a third service - `division-service` - to show how easy it is to build a new microservice and add it to integrate it with existing microservice infrastructure using our API gateway.

- Before following the steps below, clone our repository onto your local computer and make sure you are in the root directory (with the `go.mod` file)


## Creating a new RPC server (Division Service)

Create an `IDL` file e.g.<br>
- Ensure that you follow these [standards](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/thrift_idl_annotation_standards/).

```thrift
// tells thrift to generate scaffolding code in ‘Go’ in the   subdirectory kitex_gen/division/api
namespace go division.api  

// creates a new struct ‘Division request’ with the above implementation
struct DivisionRequest {
1: required string FirstNum;
2: required string SecondNum;
} 

// creates a new struct ‘Division Response’ with the above implementation
struct DivisionResponse {
1: string Quotient;
} 

// Defines the service name and the method signature of the divideNumbers method offerred by the service
service DivisionManagement {
DivisionResponse divideNumbers(1: DivisionRequest req);
} 
```

Save the `IDL` file in the `/thrift-idl` directory.

Create a new directory called `division-service` that will contain the `Kitex` code for this service. 

Then run the following command to tell `Kitex` to generate the code using your new `thrift IDL` file:

```shell
kitex -module github.com/ararchch/api-gateway -service Division ../thrift-idl/division_management.thrift

```

Update your logic in `handler.go` with the following code;

```go
package main

import (
	"context"
	"fmt"
	"strconv"

	api "github.com/ararchch/microservices/api-gateway/division-service/kitex_gen/division/api"
)

// DivisionManagementImpl implements the last service interface defined in the IDL.
type DivisionManagementImpl struct{}

// DivideNumbers implements the DivisionManagementImpl interface.
func (s *DivisionManagementImpl) DivideNumbers(ctx context.Context, req *api.DivisionRequest) (resp *api.DivisionResponse, err error) {
	
	// parse int from string of First Number
	firstNumInt, err := strconv.Atoi(req.FirstNum)
	if err != nil {
		panic(err)
	}

	// parse int from string of Second Number
	secondNumInt, err := strconv.Atoi(req.SecondNum)
	if err != nil {
		panic(err)
	}

	// divide two numbers 
	finalQuotient := firstNumInt / secondNumInt;

	// convert finalSum to string and return response of type DivisionResponse and error = nil
	return &api.DivisionResponse{
		Quotient: fmt.Sprintf("%d", finalQuotient),
	}, nil

}
```

Navigate to `main.go` in the `division-service` folder, and add the following code;

```go
package main

import (
	"log"

	api "github.com/ararchch/microservices/api-gateway/division-service/kitex_gen/division/api/divisionmanagement"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	// initate new etcd registry at port 2379
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
    if err != nil {
        log.Fatal(err)
	}
	
	// create new Kitex server for Division Service
	svr := api.NewServer(
		new(DivisionManagementImpl), // Follow DivisionManagementImpl as defined in ./handler.go
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Division"}),  // allow service to be discovered with name: "Division"
		server.WithRegistry(r), // register service on etcd registry 'r' (as declared earlier)
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
```
### Integrating Service into API Gateway

Navigate to `thrift-idl/gateway_api` and add the following code;

```go
namespace go api

struct AdditionRequest {
1: required string FirstNum (api.body="FirstNum");
2: required string SecondNum (api.body="SecondNum")
}

struct AdditionResponse {
1: string Sum;
}

struct MultiplicationRequest {
1: required string FirstNum (api.body="FirstNum");
2: required string SecondNum (api.body="SecondNum")
}

struct MultiplicationResponse {
1: string Product;
}

struct DivisionRequest {
1: required string FirstNum (api.body="FirstNum");
2: required string SecondNum (api.body="SecondNum")
}

struct DivisionResponse {
1: string Quotient;
}

service Gateway {
AdditionResponse addNumbers(1: AdditionRequest req) (api.post="/add");
MultiplicationResponse multiplyNumbers(1: MultiplicationRequest req) (api.post="/multiply");
DivisionResponse divideNumbers(1: DivisionRequest req) (api.post="/divide");
}
```

Navigate into `hertz-http-request` and run in the terminal;

```shell
hz update -idl ../thrift-idl/gateway_api.thrift
```

Navigate to `./biz/handler/api/gateway.go`

Implement `DivideNumbers` method in the file as follows; 

```go
func DivideNumbers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DivisionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// create new client (with loadbalancing, service discovery capabilities) using utils.GenerateClient feature
	divisionClient, err := utils.GenerateClient("Division")
	if err != nil {
		panic(err)
	}

	// binding req params to RPC reqest struct (following the request format declared in RPC service IDL)
	reqRpc := &divisionService.DivisionRequest{
		FirstNum:  req.FirstNum,
		SecondNum: req.SecondNum,
	}

	var respRpc api.DivisionResponse

	// calling MakeRpcRequest method declared in the utils package
	err = utils.MakeRpcRequest(ctx, divisionClient, "divideNumbers", reqRpc, &respRpc)
	if err != nil {
		panic(err)
	}

	resp := &api.DivisionResponse{
		Quotient: respRpc.Quotient,
	}

	// return to client as JSON HTTP response
	c.JSON(consts.StatusOK, resp)
}
```

Now you are done with editing the API Gateway! Thank you! **:)**
