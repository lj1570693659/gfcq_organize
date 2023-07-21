package common

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type EmployeeJobController struct {
	v1.UnimplementedEmployeeServer
}

func EmployeeJobRegister(s *grpcx.GrpcServer) {
	v1.RegisterEmployeeJobServer(s.Server, &EmployeeJobController{})
}

// GetList implements GetList
func (s *EmployeeJobController) GetList(ctx context.Context, in *v1.GetListEmployeeJobReq) (*v1.GetListEmployeeJobRes, error) {
	res, err := service.EmployeeJob().GetList(ctx, in.GetEmployeeJob(), in.GetPage(), in.GetSize())
	return res, err
}

// GetCount implements GetCount
func (s *EmployeeJobController) GetCount(ctx context.Context, in *v1.GetCountEmployeeJobReq) (*v1.GetCountEmployeeJobRes, error) {
	res, err := service.EmployeeJob().GetCount(ctx, in)
	return res, err
}

func (s *EmployeeJobController) GetOne(ctx context.Context, in *v1.GetOneEmployeeJobReq) (*v1.GetOneEmployeeJobRes, error) {
	res, err := service.EmployeeJob().GetOne(ctx, &v1.EmployeeJobInfo{
		EmployeeId: in.GetEmployeeId(),
		JobId:      in.GetJobId(),
	})
	return &v1.GetOneEmployeeJobRes{
		EmployeeJob: res,
	}, err
}

func (s *EmployeeJobController) Create(ctx context.Context, in *v1.CreateEmployeeJobReq) (*v1.CreateEmployeeJobRes, error) {
	res, err := service.EmployeeJob().Create(ctx, &v1.EmployeeJobInfo{
		EmployeeId: in.GetEmployeeId(),
		JobId:      in.GetJobId(),
		Remark:     in.GetRemark(),
	})
	return &v1.CreateEmployeeJobRes{
		EmployeeJob: res,
	}, err
}

func (s *EmployeeJobController) Modify(ctx context.Context, in *v1.ModifyEmployeeJobReq) (*v1.ModifyEmployeeJobRes, error) {
	res, err := service.EmployeeJob().Modify(ctx, &v1.EmployeeJobInfo{
		Id:         in.GetId(),
		EmployeeId: in.GetEmployeeId(),
		JobId:      in.GetJobId(),
		Remark:     in.GetRemark(),
	})
	return &v1.ModifyEmployeeJobRes{
		EmployeeJob: res,
	}, err
}

func (s *EmployeeJobController) Delete(ctx context.Context, in *v1.DeleteEmployeeJobReq) (*v1.DeleteEmployeeJobRes, error) {
	isSuccess, msg, err := service.EmployeeJob().Delete(ctx, in)
	return &v1.DeleteEmployeeJobRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
