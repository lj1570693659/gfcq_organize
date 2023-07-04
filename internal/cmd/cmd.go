package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/lj1570693659/gfcq_product/internal/controller/common"
	"github.com/lj1570693659/gfcq_product/internal/controller/user"

	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
)

var (
	// Main is the main command.
	Main = gcmd.Command{
		Name:  "gfcq_product",
		Usage: "main",
		Brief: "start grpc server of employee",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
					grpcx.Server.UnaryTracing,
					grpcx.Server.UnaryError,
					grpcx.Server.UnaryRecover,
				)}...,
			)
			s := grpcx.Server.New(c)
			user.Register(s)
			common.DepartmentRegister(s)
			common.EmployeeRegister(s)
			common.EmployeeJobRegister(s)
			common.JobRegister(s)
			common.JobLevelRegister(s)
			s.Run()
			return nil
		},
	}
)
