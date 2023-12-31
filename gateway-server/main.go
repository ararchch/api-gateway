package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {

	// initating new hertz server at port 8080
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// handle post requests containing RPC service, IDL version, and method
	h.POST("/:service/:version/:method", HandlePostRequest)

	// run the server
	h.Spin()
}
