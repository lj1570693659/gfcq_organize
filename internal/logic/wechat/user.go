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
	sWechatUser struct{}
)

func init() {
	service.RegisterWechatUser(&sWechatUser{})
}

func (s *sWechatUser) Create(ctx context.Context, in *v1.CreateWechatUserReq) (*v1.CreateWechatUserRes, error) {
	token, err := library.GetAccessToken(ctx, consts.AddressBookKey)
	fmt.Println("token------------", token)
	return nil, err
}

func (s *sWechatUser) GetOne(ctx context.Context, in *v1.GetOneWechatUserReq) (*v1.GetOneWechatUserRes, error) {
	//svr := &SWechatDepart{}
	//svr.GetList(ctx)

	return &v1.GetOneWechatUserRes{
		JobIdString: "",
	}, nil
}

// 同步部门下员工信息
func (s *sWechatUser) syncUserByDepart(ctx context.Context) error {
	// 1: 查询部门信息
	departList, err := service.Department().GetListWithoutPage(ctx, &common.DepartmentInfo{})
	if err != nil {
		return err
	}

	if len(departList.Data) > 0 {
		for _, v := range departList.Data {
			userList, err := s.getUserList(ctx, v.Id)
			if err != nil {
				return err
			}

			// 2： 同步员工信息
			if len(userList.UserList) > 0 {
				for _, uv := range userList.UserList {
					s.syncUser(ctx, uv, v.Id)
				}
			}
		}
	}
	return nil
}

func (s *sWechatUser) getUserList(ctx context.Context, departId int32) (*entity.HttpWechatUser, error) {
	tokenServer := SWechatToken{}
	token, err := tokenServer.GetToken(ctx, consts.PersonBookKey)
	if err != nil {
		return nil, err
	}

	// 1： 获取部门用户信息
	url := fmt.Sprintf(dao.GetUserListByDepart, token, departId)
	getUser, err := library.SendGetHttp(ctx, url)
	if err != nil {
		return nil, err
	}

	userList := &entity.HttpWechatUser{}
	json.Unmarshal(getUser, &userList)
	if !g.IsEmpty(userList.ErrCode) {
		return nil, errors.New(userList.ErrMsg)
	}

	return userList, err
}

func (s *sWechatUser) syncUser(ctx context.Context, userInfo entity.UserInfo, departId int32) error {
	employeeInfo := &common.CreateEmployeeReq{
		WorkNumber:     userInfo.UserId,
		UserName:       userInfo.Name,
		Phone:          userInfo.Mobile,
		Email:          userInfo.Email,
		DirectLeader:   strings.Join(userInfo.DirectLeader, ","),
		IsLeaderInDept: gconv.Int32(userInfo.IsLeader),
	}

	// 2: 获取配置信息
	//  2.1 职级
	if !g.IsEmpty(len(userInfo.Extattr.Attrs)) {
		for _, vj := range userInfo.Extattr.Attrs {
			if vj.Name == "职级" {
				jobLevelInfo, err := service.JobLevel().GetOne(ctx, &common.JobLevelInfo{Name: vj.Value})
				if err != nil && err != sql.ErrNoRows {
					return err
				}
				if err == sql.ErrNoRows {
					jobLevelInfo, err = service.JobLevel().Create(ctx, &common.JobLevelInfo{Name: vj.Value})
					if err != nil {
						return err
					}
				}
				employeeInfo.JobLevel = jobLevelInfo.Id
			}
		}
	}
	if g.IsEmpty(employeeInfo.JobLevel) {
		return errors.New("职级信息缺失")
	}

	//  2.2 岗位
	if !g.IsEmpty(userInfo.Position) {
		jobInfo, err := service.Job().GetOne(ctx, &common.JobInfo{Name: userInfo.Position, DepartId: departId})
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if err == sql.ErrNoRows || jobInfo == nil {
			jobInfo, err = service.Job().Create(ctx, &common.JobInfo{Name: userInfo.Position, DepartId: departId})
			if err != nil {
				return err
			}
		}

		employeeInfo.JobId = []int32{jobInfo.Id}

	}
	if g.IsEmpty(employeeInfo.JobId) {
		return errors.New("岗位信息缺失")
	}
	//  2.3 部门
	employeeInfo.DepartId = strings.Join(gconv.Strings(userInfo.Department), ",")
	//  2.4 状态
	employeeInfo.Status = common.StatusEnum(userInfo.Status)
	//  2.5 性别
	employeeInfo.Sex = common.SexEnum(gconv.Int32(userInfo.Gender))

	// 3: 更新员工信息
	employee, err := service.Employee().GetOne(ctx, &common.GetOneEmployeeReq{WorkNumber: userInfo.UserId})
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		createInfo, err := service.Employee().Create(ctx, employeeInfo)
		if err != nil {
			return err
		}
		employee.Employee.Id = createInfo.Employee.Id
	} else {
		_, err = service.Employee().Modify(ctx, &common.ModifyEmployeeReq{
			Id:             employee.Employee.Id,
			WorkNumber:     userInfo.UserId,
			UserName:       userInfo.Name,
			Phone:          userInfo.Mobile,
			Email:          userInfo.Email,
			JobId:          employeeInfo.JobId,
			JobLevel:       employeeInfo.JobLevel,
			DirectLeader:   strings.Join(userInfo.DirectLeader, ","),
			IsLeaderInDept: employeeInfo.IsLeaderInDept,
			DepartId:       employeeInfo.DepartId,
			Status:         common.StatusEnum(userInfo.Status),
			Sex:            employeeInfo.Sex,
		})
	}

	// 1: 更新User表
	dataInfo, err := service.User().GetInfo(ctx, entity.User{
		WorkNumber: userInfo.UserId,
	})
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if g.IsEmpty(dataInfo) || g.IsNil(dataInfo) {
		if err = service.User().Create(ctx, entity.User{
			WorkNumber: userInfo.UserId,
			EmployeeId: employee.Employee.Id,
		}); err != nil {
			return err
		}
	}
	return err
}

func (s *sWechatUser) GetList(ctx context.Context, in *v1.GetListWechatUserReq) (*v1.GetListWechatUserRes, error) {
	token, err := library.GetAccessToken(ctx, consts.AddressBookKey)
	fmt.Println("token------------", token)
	return nil, err
}
func (s *sWechatUser) GetAll(ctx context.Context, in *v1.GetAllWechatUserReq) (*v1.GetAllWechatUserRes, error) {
	token, err := library.GetAccessToken(ctx, consts.AddressBookKey)
	fmt.Println("token------------", token)
	return nil, err
}
func (s *sWechatUser) Modify(ctx context.Context, in *v1.ModifyWechatUserReq) (*v1.ModifyWechatUserRes, error) {
	token, err := library.GetAccessToken(ctx, consts.AddressBookKey)
	fmt.Println("token------------", token)
	return nil, err
}
func (s *sWechatUser) Delete(ctx context.Context, in *v1.DeleteWechatUserReq) (*v1.DeleteWechatUserRes, error) {
	token, err := library.GetAccessToken(ctx, consts.AddressBookKey)
	fmt.Println("token------------", token)
	return nil, err
}
