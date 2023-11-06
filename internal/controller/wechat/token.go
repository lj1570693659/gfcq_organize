package wechat

import (
	"context"
	"github.com/lj1570693659/gfcq_product/internal/consts"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/wechat/v1"
	"google.golang.org/grpc"
)

type WechatTokenController struct {
	v1.UnimplementedWechatUserServer
}

func WechatTokenRegister(s *grpc.Server) {
	v1.RegisterWechatUserServer(s, &WechatTokenController{})
}

// GetToken implements GetToken
func (s *WechatTokenController) GetToken(ctx context.Context) error {
	_, err := service.WechatToken().GetToken(ctx, consts.PersonBookKey)
	return err
}
