package main

import (
	management "github.com/ararchch/api-gateway/microservices/name-service/kitex_gen/name/management/namemanagement"
	"github.com/ararchch/api-gateway/utils"
)

func main() {
	servers := utils.CreateMultipleServers(
		3, // Number of servers you want to create to handle requests made to this microservice
		"Name", // name of microservice (servers will be registered under this name)
		new(NameManagementImpl), // the handler file that you defined in the previous step
		management.NewServiceInfo(), // serviceInfo file containing generated details unique to your microservice
		utils.RateLimit(1000, 1000), // optional rate limit if you wish to include it
	)

	utils.RunServers(servers) // runs all servers simultaneously

}
