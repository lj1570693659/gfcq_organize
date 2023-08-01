package common

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type JobController struct {
	v1.UnimplementedJobServer
}

func JobRegister(s *grpcx.GrpcServer) {
	v1.RegisterJobServer(s.Server, &JobController{})
}

// GetList implements GetList
func (s *JobController) GetList(ctx context.Context, in *v1.GetListJobReq) (*v1.GetListJobRes, error) {
	res, err := service.Job().GetList(ctx, in.GetJob(), in.GetPage(), in.GetSize())
	return res, err
}

func (s *JobController) GetOne(ctx context.Context, in *v1.GetOneJobReq) (*v1.GetOneJobRes, error) {
	res, err := service.Job().GetOne(ctx, &v1.JobInfo{
		Id:   in.Id,
		Name: in.GetName(),
	})
	return &v1.GetOneJobRes{
		Job: res,
	}, err
}

func (s *JobController) Create(ctx context.Context, in *v1.CreateJobReq) (*v1.CreateJobRes, error) {
	res, err := service.Job().Create(ctx, &v1.JobInfo{
		Name:     in.GetName(),
		DepartId: in.GetDepartId(),
		Remark:   in.GetRemark(),
	})
	return &v1.CreateJobRes{
		Job: res,
	}, err
}

func (s *JobController) Modify(ctx context.Context, in *v1.ModifyJobReq) (*v1.ModifyJobRes, error) {
	res, err := service.Job().Modify(ctx, &v1.JobInfo{
		Id:       in.GetId(),
		Name:     in.GetName(),
		DepartId: in.GetDepartId(),
		Remark:   in.GetRemark(),
	})
	return &v1.ModifyJobRes{
		Job: res,
	}, err
}

func (s *JobController) Delete(ctx context.Context, in *v1.DeleteJobReq) (*v1.DeleteJobRes, error) {
	isSuccess, msg, err := service.Job().Delete(ctx, in.GetId())
	return &v1.DeleteJobRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}

func (s *JobController) GetAll(ctx context.Context, in *v1.GetAllJobReq) (*v1.GetAllJobRes, error) {
	res, err := service.Job().GetAll(ctx, in.GetJob())
	return res, err
}
