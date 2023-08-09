package common

import (
	"context"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"google.golang.org/grpc"
)

type JobLevelController struct {
	v1.UnimplementedJobLevelServer
}

func JobLevelRegister(s *grpc.Server) {
	v1.RegisterJobLevelServer(s, &JobLevelController{})
}

// GetList implements GetList
func (s *JobLevelController) GetList(ctx context.Context, in *v1.GetListJobLevelReq) (*v1.GetListJobLevelRes, error) {
	res, err := service.JobLevel().GetList(ctx, in.GetJobLevel(), in.GetPage(), in.GetSize())
	return res, err
}

// GetAll implements GetAll
func (s *JobLevelController) GetAll(ctx context.Context, in *v1.GetAllJobLevelReq) (*v1.GetAllJobLevelRes, error) {
	res, err := service.JobLevel().GetAll(ctx, in.GetJobLevel(), in.GetSort())
	return res, err
}

func (s *JobLevelController) GetOne(ctx context.Context, in *v1.GetOneJobLevelReq) (*v1.GetOneJobLevelRes, error) {
	res, err := service.JobLevel().GetOne(ctx, &v1.JobLevelInfo{
		Id:   in.GetId(),
		Name: in.GetName(),
	})
	return &v1.GetOneJobLevelRes{
		JobLevel: res,
	}, err
}

func (s *JobLevelController) Create(ctx context.Context, in *v1.CreateJobLevelReq) (*v1.CreateJobLevelRes, error) {
	res, err := service.JobLevel().Create(ctx, &v1.JobLevelInfo{
		Name:   in.GetName(),
		Remark: in.GetRemark(),
	})
	return &v1.CreateJobLevelRes{
		JobLevel: res,
	}, err
}

func (s *JobLevelController) Modify(ctx context.Context, in *v1.ModifyJobLevelReq) (*v1.ModifyJobLevelRes, error) {
	res, err := service.JobLevel().Modify(ctx, &v1.JobLevelInfo{
		Id:     in.GetId(),
		Name:   in.GetName(),
		Remark: in.GetRemark(),
	})
	return &v1.ModifyJobLevelRes{
		JobLevel: res,
	}, err
}

func (s *JobLevelController) Delete(ctx context.Context, in *v1.DeleteJobLevelReq) (*v1.DeleteJobLevelRes, error) {
	isSuccess, msg, err := service.JobLevel().Delete(ctx, in.GetId())
	return &v1.DeleteJobLevelRes{
		IsSuccess: isSuccess,
		Msg:       msg,
	}, err
}
