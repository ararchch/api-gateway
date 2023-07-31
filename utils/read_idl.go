package utils

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"path/filepath"
)

func ReadIdlFromGithub(service string, version string) (string, error) {
	// creates file name using the version following pre-defined convention
	idlFileName := fmt.Sprintf("%s.thrift", version)

	// gets IDL from the [service name] branch of the IDL repo
	idlUrl := fmt.Sprintf("https://raw.githubusercontent.com/ararchch/api-gateway-idl/%s/%s", service, idlFileName)

	// Download the IDL file
	resp, err := http.Get(idlUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// extracts bytes from response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("Error in Reading IDL file")
		panic(err)
	}

	// name of new temp file to be created
	tempIdlName := fmt.Sprintf("%s-%s", service, idlFileName)
	// accesses temp directory of OS of system the server is running on
	tempDir := os.TempDir()
	// creates full local IDL path
	idlPath := filepath.Join(tempDir, tempIdlName)
	// writes temp IDL file with same data extracted from the github IDL file
	err = os.WriteFile(idlPath, body, 0644)
	if err != nil {
		fmt.Println("Error in creating temporary IDL file")
		panic(err)
	}

	return idlPath, nil
}