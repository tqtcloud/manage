package main

import (
	"flag"
	"fmt"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/internal/config"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/internal/server"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/host/rpc/types/host"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/host.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		host.RegisterHostServer(grpcServer, server.NewHostServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.DisableStat()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
