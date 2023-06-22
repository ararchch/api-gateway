package main

import (
	management "github.com/ararchch/api-gateway/mult-rpc-server/kitex_gen/multiplication/management/multiplicationmanagement"
	"log"
)

func main() {
	svr := management.NewServer(new(MultiplicationManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
