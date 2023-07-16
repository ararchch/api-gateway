package utils

import (
	"fmt"
	"net"
)

func GetUnusedPort(domain string) (*net.TCPAddr, error){
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:0", domain))
	if err != nil {
		return nil, err
	}
	
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr), nil
}