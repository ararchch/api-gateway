package main

import (
	management "github.com/ararchch/api-gateway/kitex-rpc-server/kitex_gen/addition/management/additionmanagement"
	"log"
)

func main() {
	svr := management.NewServer(new(AdditionManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
