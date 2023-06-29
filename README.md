# CloudWeGo API Gateway

An API Gateway written in `GO` that uses CloudWeGo `Kitex` and `Hertz` frameworks to develop an API Gateway that accepts `HTTP` requests encoded in `JSON` format, Utilising the Generic-Call feature of `Kitex` to translate these requests into `Thrift` binary format requests. The API Gateway will then proceed to forward the translated requests to one of the backend `RPC` servers obtained from the registry centre. 

## Tech Design

[Tech Design Link](https://drive.google.com/file/d/12YdO1ZMxGWcnbM7dVHZ8ZAMn8eJMTPbL/view?usp=drive_link) 

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

## Adding new Hertz server
There is no need to create a new `Hertz` server and we strongly recommend that you do not. **BUT** if you still want to, follow the instructions below:

Create an `IDL` file e.g.:<br>
- Ensure that you follow these [standards](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/thrift_idl_annotation_standards/).

```thrift
namespace go api

struct AdditionRequest {
    1: required string FirstNum (api.body="FirstNum");
    2: required string SecondNum (api.body="SecondNum")
}

struct AdditionResponse {
    1: string Sum;
}

service AdditionApi {
   AdditionResponse addNumbers(1: AdditionRequest req) (api.post="/add");
}
```

Save the `IDL` file in the `/thrift-idl` directory.

In the `hertz-http-server` directory, run in the terminal:

```shell
hz new -idl ../thrift-idl/[YOUR_IDL_FILE].thrift

go mod init

go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0

go mod tidy
```

Update your logic in `biz/handler/api/[YOUR_IDL_FILE].go`.

## Adding new Kitex Server 

Create an `IDL` file e.g.<br>
- Ensure that you follow these [standards](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/thrift_idl_annotation_standards/).

```thrift
namespace go addition.management

struct AdditionRequest {
    1: required string FirstNum;
    2: required string SecondNum;
}

struct AdditionResponse {
    1: string Sum;
}

service AdditionManagement {
    AdditionResponse addNumbers(1: AdditionRequest req);
}
```

Save the `IDL` file in the `/thrift-idl` directory.

In the `addition-service` directory, run in the terminal:

```shell
kitex kitex -module "your_module_name" -service "service_name" [YOUR_IDL_FILE].thrift

go mod init

go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0

go mod tidy
```

Update your logic in `handler.go`.