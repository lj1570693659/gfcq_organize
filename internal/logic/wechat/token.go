package wechat

import (
	"context"
	"fmt"
	"github.com/lj1570693659/gfcq_product/internal/library"
	"github.com/lj1570693659/gfcq_product/internal/service"
)

type (
	SWechatToken struct{}
)

func init() {
	service.RegisterWechatToken(&SWechatToken{})
}

func (s *SWechatToken) GetToken(ctx context.Context, keyName string) (string, error) {
	token, err := library.GetAccessToken(ctx, keyName)
	fmt.Println("token------------", token)
	return token, err
}
