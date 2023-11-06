package library

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lj1570693659/gfcq_product/boot"
	"github.com/lj1570693659/gfcq_product/internal/consts"
	"github.com/lj1570693659/gfcq_product/internal/dao"
	"github.com/lj1570693659/gfcq_product/internal/model/entity"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetListWithPage(query *gdb.Model, page, size int32) (*gdb.Model, int32, error) {
	totalSize, err := query.Count()
	if err != nil {
		return query, gconv.Int32(totalSize), err
	}

	query = query.Limit(gconv.Int((page-1)*size), gconv.Int(size))
	return query, gconv.Int32(totalSize), nil
}

func FindIntSlice(slice []int32, val int32) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func DeleteIntSlice(a []int32) []int32 {
	ret := make([]int32, 0, len(a))
	for _, val := range a {
		if !g.IsEmpty(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

func GetDepartRedisInfo(ctx context.Context, keyName string) (string, error) {
	value, err := boot.DepartRedis.Get(ctx, keyName).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return value, nil
}
func SetDepartRedisInfo(ctx context.Context, keyName string, value entity.WechatDepartInfo) error {
	departValue, err := boot.DepartRedis.Get(ctx, keyName).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	if g.IsEmpty(departValue) {
		valueJson, _ := json.Marshal(value)
		boot.DepartRedis.Set(ctx, keyName, string(valueJson), 0)
	}

	return nil
}

func GetAccessToken(ctx context.Context, applicationName string) (string, error) {
	token, err := boot.Redis.Get(ctx, applicationName).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	if g.IsEmpty(token) {
		accessToken, err := sendGetAccessToken(ctx, applicationName)
		if err != nil {
			return "", err
		}
		boot.Redis.Set(ctx, applicationName, accessToken.AccessToken, time.Second*3600)
		token = accessToken.AccessToken
	}
	return token, nil
}

func sendGetAccessToken(ctx context.Context, applicationName string) (*entity.WechatCryptToken, error) {
	corpId, _ := g.Config("config.yaml").Get(context.Background(), "wechat.corpId")
	corpSecret, _ := g.Config("config.yaml").Get(context.Background(), fmt.Sprintf("%s.%s.%s", "wechat", applicationName, "secret"))
	url := fmt.Sprintf(dao.GetTokenUrl, corpId, corpSecret)
	r, err := g.Client().Get(ctx, url)
	if err != nil {

	}
	defer r.Close()

	httpRead := []byte(r.ReadAllString())
	accessToken := &entity.WechatCryptToken{}
	json.Unmarshal(httpRead, &accessToken)
	return accessToken, err
}

func SendGetHttp(ctx context.Context, url string) ([]byte, error) {
	r, err := g.Client().Get(ctx, url)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return []byte(r.ReadAllString()), err
}

func SendPostHttp(ctx context.Context, url string, data interface{}) ([]byte, error) {
	r, err := g.Client().Post(ctx, url, data)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return []byte(r.ReadAllString()), err
}

func Encrypt(str string) string {
	var encryptStr string
	types, _ := g.Config("config.yaml").Get(context.Background(), "user.encrypt")
	switch types.String() {
	case consts.GSHA1:
		encryptStr = gsha1.Encrypt(str)
	case consts.MD5:
		encryptStr, _ = gmd5.Encrypt(str)
	}
	return encryptStr
}
