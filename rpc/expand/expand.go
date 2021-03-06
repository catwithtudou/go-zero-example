// Code generated by goctl. DO NOT EDIT!
// Source: expand.proto

package main

import (
	"flag"
	"fmt"
	"log"

	"shorturl/rpc/expand/internal/config"
	"shorturl/rpc/expand/internal/server"
	"shorturl/rpc/expand/internal/svc"
	expand "shorturl/rpc/expand/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rpcx"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/expand.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	expanderSrv := server.NewExpanderServer(ctx)

	s, err := rpcx.NewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		expand.RegisterExpanderServer(grpcServer, expanderSrv)
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
