package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ararchch/api-gateway/utils"
)

func main() {
	if len(os.Args) != 3 {
		panic("Please provide the service name, version as arguments")
	}

	serviceName := os.Args[1]
	version := os.Args[2]

	idlPath, err := utils.ReadIdlFromGithub(serviceName, version)
	if err != nil {
		fmt.Println("Error in reading IDL, please check IDL and retry")
		panic(err)
	}

	dirName := fmt.Sprintf("%s-service-%s", serviceName, version)

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Create a new directory path
	newDir := filepath.Join(cwd, dirName)

	// Create the new directory
	err = os.MkdirAll(newDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Change the working directory to 'serviceName-service-version'
	err = os.Chdir(newDir)
	if err != nil {
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
