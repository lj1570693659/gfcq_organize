package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"testing"
)

func Test_EmployeeJob_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewEmployeeJobClient(conn)
		res    *v1.GetListEmployeeJobRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListEmployeeJobReq{
		Page: 1,
		Size: size,
		EmployeeJob: &v1.EmployeeJobInfo{
			EmployeeId: 1,
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

func Test_EmployeeJob_GetCount(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewEmployeeJobClient(conn)
		res    *v1.GetCountEmployeeJobRes
		err    error
	)
	res, err = depert.GetCount(ctx, &v1.GetCountEmployeeJobReq{
		GroupBy:           "employee_id",
		GetFiledNameCount: "employee_id",
		EmployeeJob: &v1.EmployeeJobInfo{
			DepartId: 5,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_EmployeeJob_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewEmployeeJobClient(conn)
		res    *v1.GetOneEmployeeJobRes
		err    error
	)
	res, err = depert.GetOne(ctx, &v1.GetOneEmployeeJobReq{
		JobId: 2,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_EmployeeJob_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewEmployeeJobClient(conn)
		res    *v1.CreateEmployeeJobRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateEmployeeJobReq{
		Remark:     "13",
		EmployeeId: 2,
		JobId:      2,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_EmployeeJob_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewEmployeeJobClient(conn)
		res    *v1.ModifyEmployeeJobRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyEmployeeJobReq{
		EmployeeId: 2,
		Id:         5,
		JobId:      2,
		Remark:     "测试账号1-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_EmployeeJob_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewEmployeeJobClient(conn)
		res    *v1.DeleteEmployeeJobRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteEmployeeJobReq{
		Id: 10,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}
