package main

import (
	"log"

	management "github.com/ararchch/api-gateway/kitex-rpc-server/kitex_gen/addition/management/additionmanagement"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func main() {

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
    if err != nil {
        log.Fatal(err)
    }



	svr := management.NewServer(new(AdditionManagementImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "AdditionApi"}), server.WithRegistry(r))


	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
