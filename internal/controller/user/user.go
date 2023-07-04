package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/lj1570693659/gfcq_product/internal/dao"
	"github.com/lj1570693659/gfcq_product/internal/model/do"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/user/v1"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	// 将具体实现注册到服务对象中
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}).Insert()
	return
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	user, err := service.User().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetOneRes{
		User: user,
	}
	return
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	fmt.Println("99999===========================")
	//res = &v1.GetListRes{}
	//err = dao.User.Ctx(ctx).Page(int(req.Page), int(req.Size)).Scan(&res.Users)
	return &v1.GetListRes{}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.User().DeleteById(ctx, req.Id)
	return
}
