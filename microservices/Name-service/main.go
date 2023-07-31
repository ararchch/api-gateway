package main

import (
	management "github.com/ararchch/api-gateway/microservices/Name-service/kitex_gen/name/management/namemanagement"
	"github.com/ararchch/api-gateway/utils"
)

func main() {

	servers := utils.CreateMultipleServers(
		5, 
		"Name", 
		new(NameManagementImpl), 
		management.NewServiceInfo(), 
		utils.RateLimit(1000, 1000),
	)

	utils.RunServers(servers)
}
