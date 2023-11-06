package wechat

import (
	"context"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/wechat/v1"
	"google.golang.org/grpc"
)

type CheckInController struct {
	v1.UnimplementedWechatCheckInServer
}

func CheckInRegister(s *grpc.Server) {
	v1.RegisterWechatCheckInServer(s, &CheckInController{})
}

// GetUserCheckInDayData implements CheckInController
func (s *CheckInController) GetUserCheckInDayData(ctx context.Context, req *v1.GetUserCheckInDayDataReq) (res *v1.GetUserCheckInDayDataRes, err error) {
	return service.CheckIn().GetUserCheckInDayData(ctx, req)
}

func (s *CheckInController) SendMsg(ctx context.Context, req *v1.SendTextMsgReq) (res *v1.SendMsgRes, err error) {
	return service.CheckIn().SendMsg(ctx, req)
}
