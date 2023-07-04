package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"testing"
)

func Test_Job_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobClient(conn)
		res    *v1.GetListJobRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListJobReq{
		Page: 1,
		Size: size,
		Job: &v1.JobInfo{
			DepartId: 2,
			Name:     "项目管理",
		},
	})
	fmt.Println("res.Page=============", res.Page)
	fmt.Println("res.Size=============", res.Size)
	fmt.Println("res.TotalSize=============", res.TotalSize)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v)
	}

}

func Test_Job_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobClient(conn)
		res    *v1.GetOneJobRes
		err    error
	)
	res, err = depert.GetOne(ctx, &v1.GetOneJobReq{
		Name: "",
		Id:   1,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Job_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobClient(conn)
		res    *v1.CreateJobRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateJobReq{
		Remark:   "项目管理测试工程师",
		DepartId: 2,
		Name:     "项目管理测试工程师",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Job_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobClient(conn)
		res    *v1.ModifyJobRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyJobReq{
		Id:       3,
		Remark:   "项目管理助理工程师",
		DepartId: 2,
		Name:     "项目管理工程师",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Job_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobClient(conn)
		res    *v1.DeleteJobRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteJobReq{
		Id: 4,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}
