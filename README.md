# CloudWeGo API Gateway

An API Gateway written in `GO` that uses CloudWeGo `Kitex` and `Hertz` frameworks to develop an API Gateway that accepts `HTTP` requests encoded in `JSON` format, Utilising the Generic-Call feature of `Kitex` to translate these requests into `Thrift` binary format requests. The API Gateway will then proceed to forward the translated requests to one of the backend `RPC` servers obtained from the registry centre. 

## Tech Design

https://drive.google.com/file/d/12YdO1ZMxGWcnbM7dVHZ8ZAMn8eJMTPbL/view?usp=drive_link 

## API Gateway Design Diagram
INSERT DIAGRAM HERE 

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

Make sure `etcd` is running. Instructions [here](https://etcd.io/docs/v3.5/quickstart/) for `Mac` users. If your machine is running on `Windows`, unzip the folder and run the `etcd` application in the folder.

In the `hertz-http-server` directory, run `go run .` in the terminal.

In the `addition-service` directory, run `go run .` in the terminal.





