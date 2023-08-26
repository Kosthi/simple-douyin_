package client

import (
	"context"
	"simple-douyin/kitex_gen/message/messageservice"
	"simple-douyin/kitex_gen/user/userservice"
	"simple-douyin/pkg/constant"

	"time"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/sirupsen/logrus"
)

var UserClient userservice.Client       // interface from RPC IDL
var MessageClient messageservice.Client // interface from RPC IDL
// Init 初始化，创建rpc client
func Init(ctx context.Context) {
	var err error
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		servLog.Fatal(err)
	}
	c1, err := userservice.NewClient(
		constant.UserServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithTracer(
		// 	prometheus.NewClientTracer(
		// 		constant.UserClientTracerPort,
		// 		constant.UserClientTracerPath)), // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		servLog.Fatal(err)
	}
	UserClient = c1
	servLog.Info("User client initialized")

	c2, err := messageservice.NewClient(
		constant.MessageServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithTracer(
		// 	prometheus.NewClientTracer(
		// 		constant.MessageClientTracerPort,
		// 		constant.MessageClientTracerPath)), // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		servLog.Fatal(err)
	}

	MessageClient = c2

	servLog.Info("Message client initialized")

}
