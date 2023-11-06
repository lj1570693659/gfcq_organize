package service

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/wechat/v1"
)

type (
	ICheckIn interface {
		GetUserCheckInDayData(ctx context.Context, req *v1.GetUserCheckInDayDataReq) (*v1.GetUserCheckInDayDataRes, error)
		SendMsg(ctx context.Context, req *v1.SendTextMsgReq) (*v1.SendMsgRes, error)
	}
)

var (
	localCheckIn ICheckIn
)

func CheckIn() ICheckIn {
	if localCheckIn == nil {
		panic("implement not found for interface ICheckIn, forgot register?")
	}
	return localCheckIn
}

func RegisterCheckIn(i ICheckIn) {
	localCheckIn = i
}
