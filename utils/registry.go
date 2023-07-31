package utils

import (
    "log"

    etcd "github.com/kitex-contrib/registry-etcd"
	registry "github.com/cloudwego/kitex/pkg/registry"
)

// Single instance of ETCD registry (created upon package import) - to be used by kitex clients
var ETCDRegistry registry.Registry

// init function to create single instance above
func init() {
    var err error
    ETCDRegistry, err = etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
    if err != nil {
        log.Fatal(err)
    }
}
