package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/lj1570693659/gfcq_product/internal/cmd"
	_ "github.com/lj1570693659/gfcq_product/internal/logic/common"
)

func main() {
	//etcdLink, _ := g.Config("config.yaml").Get(context.Background(), "grpc.etcdLink")
	//grpcx.Resolver.Register(etcd.New(etcdLink.String()))
	cmd.Main.Run(gctx.New())
}
