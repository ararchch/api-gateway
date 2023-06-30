package main

import (
	"log"
	"net"

	api "github.com/ararchch/api-gateway/microservices/division-service/kitex_gen/division/api/divisionmanagement"
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
		server.WithRegistry(r), // register service on etcd registry 'r' (as declared earlier),
		server.WithServiceAddr(&net.TCPAddr{Port: 8890}),
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
