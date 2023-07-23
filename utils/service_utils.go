package utils

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/server"
)

func GetUnusedPort(domain string) (*net.TCPAddr, error){
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:0", domain))
	if err != nil {
		return nil, err
	}
	
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr), nil
}

func CreateMultipleServers(amt int, serviceName string, serviceHandler interface{}, svcInfo *serviceinfo.ServiceInfo, opts ...server.Option) [] *server.Server {
	
	servers := make([]*server.Server, amt)

	for i := 0; i < amt; i++ {

		address, err := GetUnusedPort("localhost")
		if err != nil {
			log.Fatal(err)
		}

		var options []server.Option
		options = append(options, opts...)
		options = append(options,
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
			server.WithRegistry(ETCDRegistry),
			server.WithServiceAddr(address),
		)

		

		svr := server.NewServer(options...)
		if err := svr.RegisterService(svcInfo, serviceHandler); err != nil {
			panic(err)
		}

		servers[i] = &svr
	}

	return servers
}

func RunServers(servers []*server.Server) {
	var wg sync.WaitGroup
	wg.Add(len(servers))

	for _, svr := range servers {
		go func(s *server.Server) {
			defer wg.Done()
			if err := (*s).Run(); err != nil {
				log.Printf("Server stopped with error: %v\n", err)
			}
		}(svr)
	}

	wg.Wait()
}
