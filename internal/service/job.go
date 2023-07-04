package service

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type (
	IJob interface {
		Create(ctx context.Context, info *v1.JobInfo) (*v1.JobInfo, error)
		GetOne(ctx context.Context, info *v1.JobInfo) (*v1.JobInfo, error)
		GetList(ctx context.Context, info *v1.JobInfo, page, size int32) (*v1.GetListJobRes, error)
		Modify(ctx context.Context, info *v1.JobInfo) (*v1.JobInfo, error)
		Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error)
	}
)

var (
	localJob IJob
)

func Job() IJob {
	if localJob == nil {
		panic("implement not found for interface IJob, forgot register?")
	}
	return localJob
}

func RegisterJob(i IJob) {
	localJob = i
}
