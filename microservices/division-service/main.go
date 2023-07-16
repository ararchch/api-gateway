package main

import (
	"log"

	api "github.com/ararchch/api-gateway/microservices/division-service/kitex_gen/division/api/divisionmanagement"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/ararchch/api-gateway/utils"
)

func main() {

	address, err := utils.GetUnusedPort("localhost")
	if err != nil {
		log.Fatal(err)
	}

	// create new Kitex server for Division Service
	svr := api.NewServer(
		new(DivisionManagementImpl), // Follow DivisionManagementImpl as defined in ./handler.go
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Division"}),  // allow service to be discovered with name: "Division"
		server.WithRegistry(utils.ETCDRegistry), // register service on etcd registry 'r' (as declared earlier),
		server.WithServiceAddr(address),
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
