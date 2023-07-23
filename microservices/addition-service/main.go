package main

import (
	"log"

	management "github.com/ararchch/api-gateway/microservices/addition-service/kitex_gen/addition/management/additionmanagement"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/ararchch/api-gateway/utils"
)

func main() {

	port, err := utils.GetUnusedPort("localhost")
	if err != nil {
		log.Fatal(err)
	}

	// create new Kitex server for Addition Service
	svr := management.NewServer(
		new(AdditionManagementImpl), // Follow AdditionManagementImpl as defined in ./handler.go
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Addition"}),  // allow service to be discovered with name: "Addition"
		server.WithRegistry(utils.ETCDRegistry), // register service on etcd registry 'r' (as declared earlier)
		server.WithServiceAddr(port),
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
