package common

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type EmployeeController struct {
	v1.UnimplementedEmployeeServer
}

func EmployeeRegister(s *grpcx.GrpcServer) {
	v1.RegisterEmployeeServer(s.Server, &EmployeeController{})
}

// GetList implements GetList
func (s *EmployeeController) GetList(ctx context.Context, in *v1.GetListEmployeeReq) (*v1.GetListEmployeeRes, error) {
	res, err := service.Employee().GetList(ctx, in.GetEmployee(), in.GetPage(), in.GetSize())
	return res, err
}

// GetAll implements GetAll
func (s *EmployeeController) GetAll(ctx context.Context, in *v1.GetAllEmployeeReq) (*v1.GetAllEmployeeRes, error) {
	res, err := service.Employee().GetAll(ctx, in.GetEmployee())
	return res, err
}

func (s *EmployeeController) GetOne(ctx context.Context, in *v1.GetOneEmployeeReq) (*v1.GetOneEmployeeRes, error) {
	res, err := service.Employee().GetOne(ctx, &v1.GetOneEmployeeReq{
		Id:         in.Id,
		UserName:   in.GetUserName(),
		WorkNumber: in.GetWorkNumber(),
		Phone:      in.GetPhone(),
		Email:      in.GetEmail(),
		DepartId:   in.GetDepartId(),
		JobId:      in.GetJobId(),
		JobLevel:   in.GetJobLevel(),
		Status:     in.GetStatus(),
	})
	return res, err
}

func (s *EmployeeController) Create(ctx context.Context, in *v1.CreateEmployeeReq) (*v1.CreateEmployeeRes, error) {
	res, err := service.Employee().Create(ctx, &v1.CreateEmployeeReq{
		UserName:     in.GetUserName(),
		WorkNumber:   in.GetWorkNumber(),
		Phone:        in.GetPhone(),
		Email:        in.GetEmail(),
		Sex:          in.GetSex(),
		JobLevel:     in.GetJobLevel(),
		JobId:        in.GetJobId(),
		InstructorId: in.GetInstructorId(),
		Status:       in.GetStatus(),
		Remark:       in.GetRemark(),
	})
	return res, err
}

func (s *EmployeeController) Modify(ctx context.Context, in *v1.ModifyEmployeeReq) (*v1.ModifyEmployeeRes, error) {
	res, err := service.Employee().Modify(ctx, &v1.ModifyEmployeeReq{
		Id:           in.GetId(),
		UserName:     in.GetUserName(),
		WorkNumber:   in.GetWorkNumber(),
		Phone:        in.GetPhone(),
		Email:        in.GetEmail(),
		Sex:          in.GetSex(),
		JobLevel:     in.GetJobLevel(),
		JobId:        in.GetJobId(),
		InstructorId: in.GetInstructorId(),
		Status:       in.GetStatus(),
		Remark:       in.GetRemark(),
	})
	return res, err
}

func (s *EmployeeController) Delete(ctx context.Context, in *v1.DeleteEmployeeReq) (*v1.DeleteEmployeeRes, error) {
	isSuccess, msg, err := service.Employee().Delete(ctx, in.GetId())
	return &v1.DeleteEmployeeRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
