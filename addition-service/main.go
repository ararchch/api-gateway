package main

import (
	"log"

	management "github.com/ararchch/api-gateway/addition-service/kitex_gen/addition/management/additionmanagement"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func main() {

	// initate new etcd registry at port 2379
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
    if err != nil {
        log.Fatal(err)
    }

	// create new Kitex server for Addition Service
	svr := management.NewServer(
		new(AdditionManagementImpl), // Follow AdditionManagementImpl as defined in ./handler.go
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Addition"}),  // allow service to be discovered with name: "Addition"
		server.WithRegistry(r), // register service on etcd registry 'r' (as declared earlier)
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
