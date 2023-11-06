package wechat

import (
	"context"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/wechat/v1"
	"google.golang.org/grpc"
)

type UserController struct {
	v1.UnimplementedWechatUserServer
}

func WechatUserRegister(s *grpc.Server) {
	v1.RegisterWechatUserServer(s, &UserController{})
}

// GetOne WechatUserController implements WechatUserController
func (s *UserController) GetOne(ctx context.Context, req *v1.GetOneWechatUserReq) (res *v1.GetOneWechatUserRes, err error) {
	return service.WechatUser().GetOne(ctx, req)
}
