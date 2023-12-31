package main

import (
	management "github.com/ararchch/api-gateway/microservices/Multiplication-service/kitex_gen/multiplication/management/multiplicationmanagement"
	"github.com/ararchch/api-gateway/utils"
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
