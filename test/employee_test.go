package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"testing"
)

func Test_Employee_GetList(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewEmployeeClient(conn)
		res    *v1.GetListEmployeeRes
		err    error
		size   int32 = 30
	)
	res, err = depert.GetList(ctx, &v1.GetListEmployeeReq{
		Page: 1,
		Size: size,
		Employee: &v1.EmployeeInfo{
			//DepartId: "2,7",
			JobId: "5,7",
			//Status: v1.StatusEnum_interns,
		},
	})
	fmt.Println("res.Page=============", res.Page)
	fmt.Println("res.Size=============", res.Size)
	fmt.Println("res.TotalSize=============", res.TotalSize)
	fmt.Println("Data=============", res.Data)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

	for _, v := range res.GetData() {
		fmt.Println(v.Id)
	}

}

func Test_Employee_GetOne(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewEmployeeClient(conn)
		res    *v1.GetOneEmployeeRes
		err    error
	)
	res, err = depert.GetOne(ctx, &v1.GetOneEmployeeReq{
		DepartId: []int32{},
		//UserName: "测试姣姣",
		WorkNumber: "6046053",
		Id:         0,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Employee_Create(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("organize")
		depert = v1.NewEmployeeClient(conn)
		res    *v1.CreateEmployeeRes
		err    error
	)
	res, err = depert.Create(ctx, &v1.CreateEmployeeReq{
		Remark:       "部门领导-备注信息",
		UserName:     "部门领导",
		WorkNumber:   "909903",
		Sex:          v1.SexEnum_man,
		Phone:        "18883185965",
		Email:        "18883185965@qq.com",
		JobLevel:     2,
		JobId:        []int32{8, 9},
		InstructorId: 1,
		Status:       v1.StatusEnum_interns,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Employee_Modify(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewEmployeeClient(conn)
		res    *v1.ModifyEmployeeRes
		err    error
	)
	res, err = depert.Modify(ctx, &v1.ModifyEmployeeReq{
		Id:           8,
		Remark:       "测试账号1-",
		UserName:     "测试账号1",
		WorkNumber:   "6046053",
		Sex:          v1.SexEnum_man,
		Phone:        "18883185965",
		Email:        "18883185965@qq.com",
		JobLevel:     2,
		JobId:        []int32{7, 2},
		InstructorId: 1,
		Status:       v1.StatusEnum_terminated,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}

func Test_Employee_Delete(t *testing.T) {
	var (
		ctx    = gctx.GetInitCtx()
		conn   = grpcx.Client.MustNewGrpcClientConn("employee")
		depert = v1.NewEmployeeClient(conn)
		res    *v1.DeleteEmployeeRes
		err    error
	)
	res, err = depert.Delete(ctx, &v1.DeleteEmployeeReq{
		Id: 3,
	})
	fmt.Println("res=============", res)
	fmt.Println("err=============", err)
	if err != nil {
		g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
	}

}
