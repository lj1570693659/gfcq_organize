package wechat

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lj1570693659/gfcq_product/internal/consts"
	"github.com/lj1570693659/gfcq_product/internal/dao"
	"github.com/lj1570693659/gfcq_product/internal/library"
	"github.com/lj1570693659/gfcq_product/internal/model/entity"
	"github.com/lj1570693659/gfcq_product/internal/service"
	common "github.com/lj1570693659/gfcq_protoc/common/v1"
	v1 "github.com/lj1570693659/gfcq_protoc/wechat/v1"
	"strings"
)

type (
	SWechatDepart struct{}
)

func (s *SWechatDepart) GetList(ctx context.Context) (*v1.CreateWechatUserRes, error) {
	tokenServer := SWechatToken{}
	token, err := tokenServer.GetToken(ctx, consts.PersonBookKey)
	if err != nil {
		return nil, err
	}

	// 1： 获取子部门ID列表
	url := fmt.Sprintf(dao.GetDepartmentSimpleListUrl, token, 0)
	getDepart, err := library.SendGetHttp(ctx, url)
	if err != nil {
		return nil, err
	}

	departList := &entity.HttpWechatDepart{}
	json.Unmarshal(getDepart, &departList)
	if !g.IsEmpty(departList.ErrCode) {
		return nil, errors.New(departList.ErrMsg)
	}
	if len(departList.DepartmentId) > 0 {
		for _, v := range departList.DepartmentId {
			// 2: 获取单个部门详情
			info, err := s.GetInfo(ctx, v.Id, token)
			if !g.IsEmpty(info.ErrCode) {
				return nil, errors.New(info.ErrMsg)
			}

			// 3: 缓存到Redis
			if err = library.SetDepartRedisInfo(ctx, gconv.String(v.Id), info.Department); err != nil {
				return nil, err
			}

			// 4: 更新到本地数据库
			s.syncLocalDepart(ctx, info.Department)
		}
	}
	return nil, err
}

func (s *SWechatDepart) GetInfo(ctx context.Context, id int, token string) (*entity.HttpWechatDepartInfo, error) {
	url := fmt.Sprintf(dao.GetDepartmentInfoUrl, token, id)
	getDepart, err := library.SendGetHttp(ctx, url)
	if err != nil {
		return nil, err
	}

	departInfo := &entity.HttpWechatDepartInfo{}
	json.Unmarshal(getDepart, &departInfo)
	return departInfo, err
}

func (s *SWechatDepart) syncLocalDepart(ctx context.Context, departInfo entity.WechatDepartInfo) error {
	dataInfo, err := service.Department().GetOne(ctx, &common.DepartmentInfo{Id: gconv.Int32(departInfo.Id)})
	if err != nil {
		return err
	}
	if (err != nil && err == sql.ErrNoRows) || g.IsNil(dataInfo) {
		if _, err := service.Department().Create(ctx, &common.DepartmentInfo{
			Id:               gconv.Int32(departInfo.Id),
			Name:             departInfo.Name,
			NameEn:           departInfo.NameEn,
			Pid:              gconv.Int32(departInfo.ParentId),
			DepartmentLeader: strings.Join(departInfo.DepartmentLeader, ","),
		}); err != nil {
			return err
		}
	} else {
		if _, err := service.Department().Modify(ctx, &common.DepartmentInfo{
			Id:               gconv.Int32(departInfo.Id),
			Name:             departInfo.Name,
			NameEn:           departInfo.NameEn,
			Pid:              gconv.Int32(departInfo.ParentId),
			DepartmentLeader: strings.Join(departInfo.DepartmentLeader, ","),
		}); err != nil {
			return err
		}
	}
	return nil
}
