# CloudWeGo API Gateway

An API Gateway written in `GO` that uses CloudWeGo `Kitex` and `Hertz` frameworks to develop an API Gateway that accepts `HTTP` requests encoded in `JSON` format, Utilising the Generic-Call feature of `Kitex` to translate these requests into `Thrift` binary format requests. The API Gateway will then proceed to forward the translated requests to one of the backend `RPC` servers obtained from the registry centre. 

**Note to assessors: Only very recently, we realised that we misinterpreted the instructions provided, and we made some last minute changes to our API gateway. We have created a branch called `Old-APi-Gateway` to store the previous implementation. **<br>
**Unfortunately, while we have improved our project in this Main branch, we did not have enough time to do the testing/benchmarking, so the testing done here is for the older-version. We are really sorry for the trouble.**

## Tech Design

[Tech Design Link](https://array-api-gateway.notion.site/Project-README-Milestone-3-c7df8b1dfe444b1c98e0cc0a2b53c55c) 

## API Gateway Design Diagram
![API Gateway Diagram](API-Diagram.png)

## How it works

1. Client sends HTTP request with a JSON body with a "FirstNum" and "SecondNum" attribute. <br>
2. Hertz server processes the request and sends it to the Kitex Servers. <br>
3. A generic Kitex server will be generated and registered in the `etcd` registry. <br>
4. Kitex server will process the request and convert it into `Thrift Binary Format` before sending it to the RPC servers. <br>
5. The RPC servers will process the `addition` request and returns the response. <br>
6. The API gateway will process and encode the response as an HTTP Response.

## Testing
**At the last minute, we realised that we misinterpret the instructions and improved our API gateway. Therefore these tests were done on the older-version of the project**
1. For load testing and benchmarking analysis, please refer to [Locust-ReadMe](Locust)
2. For unit testing, please refer to [Script-testing-ReadMe](Script-testing)

## Installation

Before running the project, make sure you install the following, <br>
- [Hertz](https://www.cloudwego.io/docs/hertz/getting-started/)<br>
- [Kitex](https://www.cloudwego.io/docs/kitex/getting-started/)<br>
- [etcd](https://etcd.io/docs/v3.2/install/)<br>
- [PostMan](https://www.postman.com/downloads/)<br>
- [Python](https://www.python.org/downloads/)<br>
- [Locust](https://docs.locust.io/en/stable/installation.html)

## Instructions to run the API Gateway and test using pre-implemented microservices

Make sure `etcd` is running.

- Instructions [here](https://etcd.io/docs/v3.5/quickstart/) for `Mac` users. If your machine is running on `Windows`, unzip the folder and run the `etcd` application in the folder.

In the `gateway-server` directory, run `go run .` in the terminal. (This runs the API-Gateway's HTTP server)
From inside the microservices directory, `cd` into the `Multiplication-service` directory. Then run the command `go run .`
(You can choose to follow the same steps to run the `Name-service` which is also located in the microservices directory)

Test with `Postman` using the following post request: 

**Address**: `http://127.0.0.1:8080/Multiplication/v1/multiplyNumbers` 

with the JSON body of:

```json
{
  "FirstNum" : "1",
  "SecondNum" : "2"
}
```

Anticipate a response of:

```json
{
  "Product": "2"
}
```

Similarly you can test the `Name` service using the following post request details:
**Address**: `http://127.0.0.1:8080/Name/v1/helloName` 

with the JSON body of:

```json
{
    "Name": "Winne The Pooh"
}
```

Anticipate a response of:

```json
{
    "HelloResp": "Hello, Winne The Pooh"
}
```
That is all for testing the service on the gateway! 

## Creating a new Divison microservice and integrating it on the API Gateway

To create a new service from scratch, you must first create its IDL file. For our gateway, this is written in Apache Thrift. 
For this demo, we will be creating a new service called the `Addition` service. The IDL file for the `Addition` service is as follows:
```
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
Once you have your IDL file ready, head over to our IDL storage repo: github.com/ararchch/api-gateway-idl

Create a new branch that branches from `main`. Name this branch with your service's name. Make sure that this is the service name you use from this point on, particularly for service registry in the future. If there is a mismatch, the gateway will not work as intended as this name is used to map the service to its IDL via github branches. 

Upload your IDL file into this branch. This gateway supports versioning so you can upload multiple versions of the IDL over time. We recommend you follow the format of naming thrift files `v1.thrift`, `v2.thrift` and so on, however, you can change the version name if you wish, just ensure that you use it consistently throughout the rest of the gateway implementation. In our case, this has already been done for you with the `v1.thrift` file in the `Addition` branch.

Once you have uploaded your IDL to your branch, it should be accessible via `https://raw.githubusercontent.com/ararchch/api-gateway-idl/[service]/[version].thrift`
Do check that it works, before proceeding. Eg: `https://raw.githubusercontent.com/ararchch/api-gateway-idl/Addition/v1.thrift`

Once you are ready, `cd` into the microservices directory, and run the generate_service.go program using the following command:
`go run generate_service.go [service] [version]` In our demo, this would be `go run generate_service.go Addition v1`.

This should generate a directory within the microservices directory, called `[ServiceName]-service-[version]`. Eg: `Addition-service-v1`

**You may notice that the pre-implemented services do not follow this convention. This is because they were generated before we realised that a potential clash may arise if there are multiple versions of IDLs used to generate the same service. As such we changed the directory naming convention to compensate for any future generated directories, however, the Multiplcation and Name services still follow the previous implementation.**

`cd` into this generated directory, and open the `handler.go` file. This is where you will enter the logic for the microservice. In the case of Addition, add the following code to the `AddNumbers` method:

```
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

// add two numbers together
finalSum := firstNumInt + secondNumInt

// convert finalSum to string and return response of type AdditionResponse and error = nil
return &management.AdditionResponse{
	Sum: fmt.Sprintf("%d", finalSum),
}, nil
```
Remember to tidy your imports after this. Then proceed to the `main.go` file and add the following code to generate and run servers that will run your microservice:

```
servers := utils.CreateMultipleServers(
		3, // Number of servers you want to create to handle requests made to this microservice
		"Addition", // name of microservice (servers will be registered under this name). **Make sure this is the same as the name of your branch in the idl repo**
		new(AdditionManagementImpl), // the handler file that you defined in the previous step
		management.NewServiceInfo(), // serviceInfo file containing generated details unique to your microservice
		utils.RateLimit(1000, 1000), // optional rate limiter
	)

utils.RunServers(servers)
```
Note that this uses a number of useful methods offerred by the utils package within the gateway. You do not have to use them if you don't want to though. In addition kitex offers a number of more options for the gateway which have not been wrapped by the current utils package. You can feel free to add any in to the CreateMultipleServers method if you wish (after the RateLimit option)

Once you have done this, you have finished implementing your service. To run and test, follow the instructions indicated in the 'pre-implemented microservices' section. Remember to run the gateway-server if you havent yet. If it is already running, this is not necessary as this gateway is dynamic. 

For our demo, (and as usual) run the gateway by cd-ing into the `gateway-server` directory, and running the `go run .` command.
You can run the newly created Addition service by changing to its directory and running `go run .` as well. 
Finally, you can test by sending a post request as follows:

**Address**: `http://127.0.0.1:8080/Addition/v1/addNumbers` 

with the JSON body of:

```json
{
  "FirstNum" : "1",
  "SecondNum" : "2"
}
```

Anticipate a response of:

```json
{
  "Sum": "3"
}
```
You have now fully utilised the gateway!
