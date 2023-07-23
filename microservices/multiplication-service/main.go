package main

import (
	//"fmt"
	// "log"

	management "github.com/ararchch/api-gateway/microservices/multiplication-service/kitex_gen/multiplication/management/multiplicationmanagement"
	"github.com/ararchch/api-gateway/utils"
	//"github.com/cloudwego/kitex/pkg/limit"
	// "github.com/cloudwego/kitex/pkg/rpcinfo"
	//"github.com/cloudwego/kitex/server"
)

func main() {

	servers := utils.CreateMultipleServers(
		3, 
		"Multiplication",
		new(MultiplicationManagementImpl), 
		management.NewServiceInfo(),
		utils.RateLimit(1000, 1000),
	)

	utils.RunServers(servers)

}
