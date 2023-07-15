package utils

import (
    "log"

    etcd "github.com/kitex-contrib/registry-etcd"
	registry "github.com/cloudwego/kitex/pkg/registry"
)

var ETCDRegistry registry.Registry

func init() {
    var err error
    ETCDRegistry, err = etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
    if err != nil {
        log.Fatal(err)
    }
}
