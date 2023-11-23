package test

import (
	"context"
	"fmt"
	v1 "github.com/lj1570693659/gfcq_protoc/wechat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func Test_Wechat_Checkin(t *testing.T) {
	OrganizeServer, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	conn := v1.NewWechatCheckInClient(OrganizeServer)
	res, err := conn.SendMsg(context.Background(), &v1.SendTextMsgReq{
		Touser:  []string{"6046051"},
		Msgtype: "text",
		Content: &v1.TextContent{
			Content: "jiaojiao",
		},
	})
	fmt.Println("conn=============", conn)
	fmt.Println("res=============", res)
}
