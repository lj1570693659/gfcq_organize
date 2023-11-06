package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"testing"
)

func Test_Depart_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewDepartmentClient(conn)
		res    *v1.GetListDepartmentRes
		err    error
		size   int32 = 3
	)
	res, err = depert.GetList(ctx, &v1.GetListDepartmentReq{
		Page: 2,
		Size: size,
	})
	fmt.Println("res=============", res.TotalSize)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v)
	}

}

func Test_Depart_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("jiao")
		depert = v1.NewDepartmentClient(conn)
		res    *v1.GetOneDepartmentRes
		err    error
	)
	res, err = depert.GetOne(ctx, &v1.GetOneDepartmentReq{
		Name: "总经办",
		Id:   0,
	})
	fmt.Println("conn=============", conn)
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Depart_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewDepartmentClient(conn)
		res    *v1.CreateDepartmentRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateDepartmentReq{
		Name:   "安环室",
		Pid:    4,
		Remark: "安环室",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Depart_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewDepartmentClient(conn)
		res    *v1.ModifyDepartmentRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyDepartmentReq{
		Id:     6,
		Name:   "行政室",
		Pid:    40,
		Remark: "行政室-备注信息",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Depart_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewDepartmentClient(conn)
		res    *v1.DeleteDepartmentRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteDepartmentReq{
		Id: 7,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Depart_GetAll(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewDepartmentClient(conn)
		res    *v1.GetListWithoutDepartmentRes
		err    error
	)
	res, err = depert.GetListWithoutPage(ctx, &v1.GetListWithoutDepartmentReq{
		Department: &v1.DepartmentInfo{
			Pid: 4,
		},
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}
