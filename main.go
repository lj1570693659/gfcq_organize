package main

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	_ "github.com/lj1570693659/gfcq_product/boot"
	"github.com/lj1570693659/gfcq_product/internal/controller/common"
	"github.com/lj1570693659/gfcq_product/internal/controller/wechat"
	_ "github.com/lj1570693659/gfcq_product/internal/logic/common"
	_ "github.com/lj1570693659/gfcq_product/internal/logic/user"
	_ "github.com/lj1570693659/gfcq_product/internal/logic/wechat"
	"google.golang.org/grpc"
	"net"
)

func main() {
	etcdLink, err := g.Config("config.yaml").Get(context.Background(), "grpc.etcdLink")

	//3. 设置监听， 指定 IP、port
	listener, err := net.Listen("tcp", etcdLink.String())
	if err != nil {
		fmt.Println(err)
	}

	//1. 初始一个 grpc 对象
	grpcServer := grpc.NewServer()
	//2. 注册服务
	common.DepartmentRegister(grpcServer)
	common.EmployeeRegister(grpcServer)
	common.EmployeeJobRegister(grpcServer)
	common.JobRegister(grpcServer)
	common.JobLevelRegister(grpcServer)
	wechat.WechatUserRegister(grpcServer)
	wechat.CheckInRegister(grpcServer)

	// 4、退出关闭监听
	defer listener.Close()
	//5、启动服务
	grpcServer.Serve(listener)

	//s := grpcx.Server.New()
	//wechat.WechatTokenRegister(s)
	//wechat.WechatUserRegister(s)
	//s.Run()
}
