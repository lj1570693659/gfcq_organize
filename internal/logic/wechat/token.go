package wechat

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
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

func (s *SWechatToken) GetAgentId(ctx context.Context, keyName string) (int, error) {
	agentId, err := g.Config("config.yaml").Get(context.Background(), fmt.Sprintf("%s.%s.%s", "wechat", keyName, "agentId"))
	fmt.Println("token------------", agentId)
	return agentId.Int(), err
}
