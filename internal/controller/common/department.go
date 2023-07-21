package common

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type DepartmentController struct {
	v1.UnimplementedDepartmentServer
}

func DepartmentRegister(s *grpcx.GrpcServer) {
	v1.RegisterDepartmentServer(s.Server, &DepartmentController{})
}

// GetList implements GetList
func (s *DepartmentController) GetList(ctx context.Context, in *v1.GetListDepartmentReq) (*v1.GetListDepartmentRes, error) {
	res, err := service.Department().GetList(ctx, in.GetDepartment(), in.GetPage(), in.GetSize())
	return res, err
}

func (s *DepartmentController) GetListWithoutPage(ctx context.Context, in *v1.GetListWithoutDepartmentReq) (*v1.GetListWithoutDepartmentRes, error) {
	res, err := service.Department().GetListWithoutPage(ctx, in.GetDepartment())
	return res, err
}

func (s *DepartmentController) GetOne(ctx context.Context, in *v1.GetOneDepartmentReq) (*v1.GetOneDepartmentRes, error) {
	res, err := service.Department().GetOne(ctx, &v1.DepartmentInfo{
		Id:   in.Id,
		Name: in.GetName(),
	})
	return &v1.GetOneDepartmentRes{
		Department: res,
	}, err
}

func (s *DepartmentController) Create(ctx context.Context, in *v1.CreateDepartmentReq) (*v1.CreateDepartmentRes, error) {
	res, err := service.Department().Create(ctx, &v1.DepartmentInfo{
		Name:   in.GetName(),
		Pid:    in.GetPid(),
		Remark: in.GetRemark(),
	})
	return &v1.CreateDepartmentRes{
		Department: res,
	}, err
}

func (s *DepartmentController) Modify(ctx context.Context, in *v1.ModifyDepartmentReq) (*v1.ModifyDepartmentRes, error) {
	res, err := service.Department().Modify(ctx, &v1.DepartmentInfo{
		Id:     in.GetId(),
		Name:   in.GetName(),
		Pid:    in.GetPid(),
		Remark: in.GetRemark(),
	})
	return &v1.ModifyDepartmentRes{
		Department: res,
	}, err
}

func (s *DepartmentController) Delete(ctx context.Context, in *v1.DeleteDepartmentReq) (*v1.DeleteDepartmentRes, error) {
	isSuccess, msg, err := service.Department().Delete(ctx, in.GetId())
	return &v1.DeleteDepartmentRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
