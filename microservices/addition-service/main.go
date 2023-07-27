package main

import (
	management "github.com/ararchch/api-gateway/microservices/addition-service/kitex_gen/addition/management/additionmanagement"
	"github.com/ararchch/api-gateway/utils"
)

func main() {

	servers := utils.CreateMultipleServers(
		3, 
		"Addition",
		new(AdditionManagementImpl), 
		management.NewServiceInfo(),
		utils.RateLimit(1000, 1000),
	)

	utils.RunServers(servers)
}
