package service

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type (
	IDepartment interface {
		Create(ctx context.Context, info *v1.DepartmentInfo) (*v1.DepartmentInfo, error)
		GetOne(ctx context.Context, info *v1.DepartmentInfo) (*v1.DepartmentInfo, error)
		GetList(ctx context.Context, info *v1.DepartmentInfo, page, size int32) (*v1.GetListDepartmentRes, error)
		Modify(ctx context.Context, info *v1.DepartmentInfo) (*v1.DepartmentInfo, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localDepartment IDepartment
)

func Department() IDepartment {
	if localDepartment == nil {
		panic("implement not found for interface IDepartment, forgot register?")
	}
	return localDepartment
}

func RegisterDepartment(i IDepartment) {
	localDepartment = i
}
