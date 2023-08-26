// Code generated by hertz generator.

package main

import (
	"simple-douyin/api/biz/client"
	"simple-douyin/api/biz/handler"
	"simple-douyin/api/biz/middleware"
	"simple-douyin/pkg/constant"

	"github.com/cloudwego/hertz/pkg/app/server"
	prometheus "github.com/hertz-contrib/monitor-prometheus"
)

func Init() {
	// init rpc client
	client.InitClient()
	// init jwt middleware
	middleware.InitJwt()
}

func main() {
	// for non-framework part
	Init()
	h := server.New(
		server.WithMaxRequestBodySize(constant.MaxVideoSize),
		server.WithTracer(prometheus.NewServerTracer(":9101", "/hertz")),
	)
	h.NoRoute(handler.Default)
	h.NoMethod(handler.Default)
	register(h)
	h.Spin()
}
