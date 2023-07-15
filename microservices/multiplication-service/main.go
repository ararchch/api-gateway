package main

import (
	"log"

	"github.com/ararchch/api-gateway/utils"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	management "github.com/ararchch/api-gateway/microservices/multiplication-service/kitex_gen/multiplication/management/multiplicationmanagement"

)

func main() {
	// create new Kitex server for Multiplication Service
	svr := management.NewServer(
		new(MultiplicationManagementImpl), // Follow MultiplicationManagementImpl as defined in ./handler.go
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Multiplication"}),  // allow service to be discovered with name: "Multiplication"
		server.WithRegistry(utils.ETCDRegistry), // register service on etcd registry 'r' (as declared earlier)
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
