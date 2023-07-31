package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ararchch/api-gateway/utils"
)

func main() {
	if len(os.Args) != 2 {
		panic("Please provide the service name, version as arguments")
	}

	serviceName := os.Args[0]
	version := os.Args[1]

	idlPath, err := utils.ReadIdlFromGithub(serviceName, version)
	if err != nil {
		fmt.Println("Error in reading IDL, please check IDL and retry")
		panic(err)
	}

	// Run the kitex command
	cmd := exec.Command("kitex", "-module", "github.com/ararchch/api-gateway", "-service", serviceName, idlPath)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// Delete the temporary file
	err = os.Remove(idlPath)
	if err != nil {
		panic(err)
	}
}
