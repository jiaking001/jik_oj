package main

import (
	"app/rpc/question/model"
	"flag"
	"fmt"
	"github.com/joho/godotenv"

	"app/rpc/question/internal/config"
	"app/rpc/question/internal/server"
	"app/rpc/question/internal/svc"
	"app/rpc/question/types/question"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/question/etc/question.yaml", "the config file")

func main() {
	// 加载环境变量
	_ = godotenv.Load("rpc/question/.env")
	// 初始化数据库
	model.Init()
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		question.RegisterQuestionServer(grpcServer, server.NewQuestionServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
