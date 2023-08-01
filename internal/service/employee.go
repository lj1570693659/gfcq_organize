package service

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type (
	IEmployee interface {
		Create(ctx context.Context, info *v1.CreateEmployeeReq) (*v1.CreateEmployeeRes, error)
		GetOne(ctx context.Context, info *v1.GetOneEmployeeReq) (*v1.GetOneEmployeeRes, error)
		GetList(ctx context.Context, info *v1.EmployeeInfo, page, size int32) (*v1.GetListEmployeeRes, error)
		GetAll(ctx context.Context, info *v1.EmployeeInfo) (*v1.GetAllEmployeeRes, error)
		Modify(ctx context.Context, info *v1.ModifyEmployeeReq) (*v1.ModifyEmployeeRes, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localEmployee IEmployee
)

func Employee() IEmployee {
	if localEmployee == nil {
		panic("implement not found for interface IEmployee, forgot register?")
	}
	return localEmployee
}

func RegisterEmployee(i IEmployee) {
	localEmployee = i
}
