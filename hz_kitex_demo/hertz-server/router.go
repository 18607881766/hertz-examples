// Code generated by hertz generator.

package main

import (
	handler "hertz-examples/hz_demo/hertz-server/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// Define routes outside of idl
	r.GET("/ping2", handler.Ping)
	// your code ...
}