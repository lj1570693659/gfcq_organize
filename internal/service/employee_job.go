package service

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type (
	IEmployeeJob interface {
		Create(ctx context.Context, info *v1.EmployeeJobInfo) (*v1.EmployeeJobInfo, error)
		GetOne(ctx context.Context, info *v1.EmployeeJobInfo) (*v1.EmployeeJobInfo, error)
		GetList(ctx context.Context, info *v1.EmployeeJobInfo, page, size int32) (*v1.GetListEmployeeJobRes, error)
		GetCount(ctx context.Context, info *v1.GetCountEmployeeJobReq) (*v1.GetCountEmployeeJobRes, error)
		Modify(ctx context.Context, info *v1.EmployeeJobInfo) (*v1.EmployeeJobInfo, error)
		Delete(ctx context.Context, info *v1.DeleteEmployeeJobReq) (isSuccess bool, msg string, err error)
	}
)

var (
	localEmployeeJob IEmployeeJob
)

func EmployeeJob() IEmployeeJob {
	if localEmployeeJob == nil {
		panic("implement not found for interface IEmployeeJob, forgot register?")
	}
	return localEmployeeJob
}

func RegisterEmployeeJob(i IEmployeeJob) {
	localEmployeeJob = i
}
