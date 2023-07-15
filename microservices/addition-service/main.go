package main

import (
	"log"
	"net"

	management "github.com/ararchch/api-gateway/microservices/addition-service/kitex_gen/addition/management/additionmanagement"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/ararchch/api-gateway/utils"
)

func main() {

	// create new Kitex server for Addition Service
	svr := management.NewServer(
		new(AdditionManagementImpl), // Follow AdditionManagementImpl as defined in ./handler.go
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Addition"}),  // allow service to be discovered with name: "Addition"
		server.WithRegistry(utils.ETCDRegistry), // register service on etcd registry 'r' (as declared earlier)
		server.WithServiceAddr(&net.TCPAddr{Port: 8889}),
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
