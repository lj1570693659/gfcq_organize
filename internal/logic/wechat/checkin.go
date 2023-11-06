package wechat

import (
	"context"
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
	"strconv"
	"time"
)

type (
	sCheckIn struct{}
)

func init() {
	service.RegisterCheckIn(&sCheckIn{})
}

func (s *sCheckIn) GetUserCheckInDayData(ctx context.Context, req *v1.GetUserCheckInDayDataReq) (*v1.GetUserCheckInDayDataRes, error) {
	res := &v1.GetUserCheckInDayDataRes{}
	resData := &entity.HttpWechatCheckInDayDataRes{}
	// 1: 查询部门信息
	if req.DepartId > 0 {
		departList, err := service.Department().GetListWithoutPage(ctx, &common.DepartmentInfo{Id: req.DepartId})
		if err != nil {
			return res, err
		}
		if len(departList.Data) > 0 {
			for _, v := range departList.Data {
				employList, err := service.Employee().GetAll(ctx, &common.EmployeeInfo{DepartId: gconv.String(v.Id)})
				if err != nil {
					return res, err
				}
				if len(employList.Data) > 0 {
					for _, ev := range employList.Data {
						req.WorkNumber = append(req.WorkNumber, ev.WorkNumber)
					}
				}
			}
		}
	}

	if len(req.WorkNumber) > 0 {
		resData, _ = s.getUserCheckInDayData(ctx, req)
		if len(resData.Datas) > 0 {
			for _, v := range resData.Datas {
				otDuration, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", gconv.Float64(v.OtInfo.OtDuration)/3600), 64)
				startHour := ""
				startMinute := ""
				endHour := ""
				endMinute := ""
				if v.SummaryInfo.EarliestTime > 0 {
					startHour = fmt.Sprintf("%02d", v.SummaryInfo.EarliestTime/3600)
					startMinute = fmt.Sprintf("%02d", (v.SummaryInfo.EarliestTime%3600)/60)
				}
				if v.SummaryInfo.LastestTime > 0 {
					endHour = fmt.Sprintf("%02d", v.SummaryInfo.LastestTime/3600)
					endMinute = fmt.Sprintf("%02d", (v.SummaryInfo.LastestTime%3600)/60)
				}
				res.Data = append(res.Data, &v1.GetUserCheckInDayData{
					SummaryInfo: &v1.SummaryInfo{
						RecordType:      v.SummaryInfo.RecordType,
						CheckinCount:    v.SummaryInfo.CheckinCount,
						RegularWorkSec:  v.SummaryInfo.RegularWorkSec,
						StandardWorkSec: v.SummaryInfo.StandardWorkSec,
						EarliestTime:    fmt.Sprintf("%s:%s", startHour, startMinute),
						LastestTime:     fmt.Sprintf("%s:%s", endHour, endMinute),
					},
					OtInfo: &v1.OtInfo{
						OtStatus:              gconv.Int32(v.OtInfo.OtStatus),
						OtDuration:            gconv.Float32(otDuration),
						ExceptionDuration:     gconv.Int32s(v.OtInfo.ExceptionDuration),
						WorkdayOverAsVacation: gconv.Int32(v.OtInfo.WorkdayOverAsVacation),
						WorkdayOverAsMoney:    gconv.Int32(v.OtInfo.WorkdayOverAsMoney),
					},
					BaseInfo: &v1.BaseInfo{
						Date:        time.Unix(gconv.Int64(v.BaseInfo.Date), 0).Format("2006-01-02"),
						RecordType:  gconv.Int32(v.BaseInfo.RecordType),
						Name:        v.BaseInfo.Name,
						NameEx:      v.BaseInfo.NameEx,
						DepartsName: v.BaseInfo.DepartsName,
						Acctid:      v.BaseInfo.Acctid,
						DayType:     gconv.Int32(v.BaseInfo.DayType),
						RuleInfo: &v1.RuleInfo{
							Groupid:      gconv.Int32(v.BaseInfo.RuleInfo.Groupid),
							Groupname:    v.BaseInfo.RuleInfo.Groupname,
							Scheduleid:   gconv.Int32(v.BaseInfo.RuleInfo.Scheduleid),
							Schedulename: v.BaseInfo.RuleInfo.Schedulename,
						},
					},
				})
			}
		}
	}
	return res, nil
}

func (s *sCheckIn) getUserCheckInDayData(ctx context.Context, req *v1.GetUserCheckInDayDataReq) (*entity.HttpWechatCheckInDayDataRes, error) {
	res := &entity.HttpWechatCheckInDayDataRes{}
	tokenServer := SWechatToken{}
	token, err := tokenServer.GetToken(ctx, consts.CheckIn)
	if err != nil {
		return res, err
	}

	// 1： 获取部门用户信息
	url := fmt.Sprintf(dao.GetUserCheckInDayData, token)
	reqData := entity.HttpWechatCheckInDayDataReq{
		UseridList: req.WorkNumber,
		StartTime:  gconv.Uint32(req.StartTime),
		EndTime:    gconv.Uint32(req.EndTime),
	}
	reqDataByte, _ := json.Marshal(reqData)
	getDayData, err := library.SendPostHttp(ctx, url, string(reqDataByte))
	if err != nil {
		return res, err
	}

	entityData := &entity.HttpWechatCheckInDayDataRes{}
	json.Unmarshal(getDayData, &entityData)
	if !g.IsEmpty(entityData.ErrCode) {
		return res, errors.New(entityData.ErrMsg)
	}
	res.Datas = entityData.Datas
	return res, nil
}

func (s *sCheckIn) SendMsg(ctx context.Context, req *v1.SendTextMsgReq) (*v1.SendMsgRes, error) {
	res := &v1.SendMsgRes{}
	//tokenServer := SWechatToken{}
	//token, err := tokenServer.GetToken(ctx, consts.CheckIn)
	//if err != nil {
	//	return res, err
	//}
	//
	//// 1： 获取部门用户信息
	//url := fmt.Sprintf(dao.SendMsgApi, token)
	//reqData := entity.SendTextMsgApiReq{
	//	Touser
	//	Toparty
	//	Totag
	//	Msgtype
	//	Agentid
	//	Text
	//}
	//reqDataByte, _ := json.Marshal(reqData)
	//getDayData, err := library.SendPostHttp(ctx, url, string(reqDataByte))
	//if err != nil {
	//	return res, err
	//}
	return res, nil
}
