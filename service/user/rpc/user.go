package main

import (
	"flag"
	"fmt"
	"github.com/tqtcloud/manage/common/interceptor/rpcserver"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/tqtcloud/manage/service/user/rpc/internal/config"
	"github.com/tqtcloud/manage/service/user/rpc/internal/server"
	"github.com/tqtcloud/manage/service/user/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 关闭输出的统计日志(stat)
	logx.DisableStat()
	//rpc log,grpc的全局拦截器
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
