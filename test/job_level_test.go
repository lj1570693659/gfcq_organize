package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"testing"
)

func Test_JobLevel_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("jiao")
		depert = v1.NewJobLevelClient(conn)
		res    *v1.GetListJobLevelRes
		err    error
		size   int32 = 3
	)
	fmt.Println("conn=============", conn)
	res, err = depert.GetList(ctx, &v1.GetListJobLevelReq{
		Page: 1,
		Size: size,
		JobLevel: &v1.JobLevelInfo{
			Name: "11",
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

func Test_JobLevel_GetAll(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewJobLevelClient(conn)
		res    *v1.GetAllJobLevelRes
		err    error
	)
	res, err = depert.GetAll(ctx, &v1.GetAllJobLevelReq{
		JobLevel: &v1.JobLevelInfo{},
		Sort:     v1.OrderEnum_desc,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v)
	}

}

func Test_JobLevel_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobLevelClient(conn)
		res    *v1.GetOneJobLevelRes
		err    error
	)
	res, err = depert.GetOne(ctx, &v1.GetOneJobLevelReq{
		Name: "12",
		Id:   0,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_JobLevel_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobLevelClient(conn)
		res    *v1.CreateJobLevelRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateJobLevelReq{
		Remark: "13",
		Name:   "13",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_JobLevel_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobLevelClient(conn)
		res    *v1.ModifyJobLevelRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyJobLevelReq{
		Id:     3,
		Remark: "测试账号1-备注信息",
		Name:   "14",
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_JobLevel_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewJobLevelClient(conn)
		res    *v1.DeleteJobLevelRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteJobLevelReq{
		Id: 1,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}
