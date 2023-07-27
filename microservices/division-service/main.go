package main

import (
	api "github.com/ararchch/api-gateway/microservices/division-service/kitex_gen/division/api/divisionmanagement"
	"github.com/ararchch/api-gateway/utils"
)

func main() {

	servers := utils.CreateMultipleServers(
		3, 
		"Division",
		new(DivisionManagementImpl), 
		api.NewServiceInfo(),
		utils.RateLimit(1000, 1000),
	)

	utils.RunServers(servers)
}
