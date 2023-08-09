package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product/internal/controller/common"
	"google.golang.org/grpc"
	"net"
)

func main() {
	etcdLink, _ := g.Config("config.yaml").Get(context.Background(), "grpc.etcdLink")
	fmt.Println(etcdLink.String())
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

	// 4退出关闭监听
	defer listener.Close()
	//5、启动服务
	grpcServer.Serve(listener)
}
