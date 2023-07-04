package service

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type (
	IJobLevel interface {
		Create(ctx context.Context, info *v1.JobLevelInfo) (*v1.JobLevelInfo, error)
		GetOne(ctx context.Context, info *v1.JobLevelInfo) (*v1.JobLevelInfo, error)
		GetList(ctx context.Context, info *v1.JobLevelInfo, page, size int32) (*v1.GetListJobLevelRes, error)
		Modify(ctx context.Context, info *v1.JobLevelInfo) (*v1.JobLevelInfo, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localJobLevel IJobLevel
)

func JobLevel() IJobLevel {
	if localDepartment == nil {
		panic("implement not found for interface IJobLevel, forgot register?")
	}
	return localJobLevel
}

func RegisterJobLevel(i IJobLevel) {
	localJobLevel = i
}
