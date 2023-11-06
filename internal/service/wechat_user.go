package service

import (
	"context"
	v1 "github.com/lj1570693659/gfcq_protoc/wechat/v1"
)

type (
	IWechatUser interface {
		Create(ctx context.Context, info *v1.CreateWechatUserReq) (*v1.CreateWechatUserRes, error)
		GetOne(ctx context.Context, info *v1.GetOneWechatUserReq) (*v1.GetOneWechatUserRes, error)
		GetList(ctx context.Context, info *v1.GetListWechatUserReq) (*v1.GetListWechatUserRes, error)
		Modify(ctx context.Context, info *v1.ModifyWechatUserReq) (*v1.ModifyWechatUserRes, error)
		Delete(ctx context.Context, info *v1.DeleteWechatUserReq) (res *v1.DeleteWechatUserRes, err error)
	}
)

var (
	localWechatUser IWechatUser
)

func WechatUser() IWechatUser {
	if localWechatUser == nil {
		panic("implement not found for interface IWechatUser, forgot register?")
	}
	return localWechatUser
}

func RegisterWechatUser(i IWechatUser) {
	localWechatUser = i
}
