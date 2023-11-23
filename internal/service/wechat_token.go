package service

import (
	"context"
)

type (
	IWechatToken interface {
		GetToken(ctx context.Context, keyName string) (string, error)
		GetAgentId(ctx context.Context, keyName string) (int, error)
		//DeleteById(ctx context.Context, uid uint64) error
	}
)

var (
	localWechatToken IWechatToken
)

func WechatToken() IWechatToken {
	if localWechatToken == nil {
		panic("implement not found for interface IWechatToken, forgot register?")
	}
	return localWechatToken
}

func RegisterWechatToken(i IWechatToken) {
	localWechatToken = i
}
