package common

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lj1570693659/gfcq_product/internal/dao"
	"github.com/lj1570693659/gfcq_product/internal/library"
	"github.com/lj1570693659/gfcq_product/internal/model/do"
	"github.com/lj1570693659/gfcq_product/internal/model/entity"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type (
	sJobLevel struct{}
)

func init() {
	service.RegisterJobLevel(&sJobLevel{})
}

func (s *sJobLevel) Create(ctx context.Context, in *v1.JobLevelInfo) (*v1.JobLevelInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("职级名称不能为空")
	}

	// 不能重名
	info, err := s.GetOne(ctx, &v1.JobLevelInfo{
		Name: in.GetName(),
	})
	if (err != nil && err != sql.ErrNoRows) || !g.IsNil(info) {
		return in, err
	}
	if !g.IsNil(info) && info.Id > 0 {
		return in, errors.New("职级已存在，请重新命名")
	}

	data := do.JobLevel{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return in, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.JobLevel.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return in, err
	}
	in.Id = gconv.Int32(lastInsertId)
	return in, nil
}
func (s *sJobLevel) GetOne(ctx context.Context, in *v1.JobLevelInfo) (*v1.JobLevelInfo, error) {
	var jobLevel *v1.JobLevelInfo
	query := dao.JobLevel.Ctx(ctx)

	if len(in.GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.JobLevel.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetName(), "%")})
	}
	if in.GetId() > 0 {
		query = query.Where(dao.JobLevel.Columns().Id, in.GetId())
	}

	err := query.Scan(&jobLevel)

	return jobLevel, err
}

func (s *sJobLevel) GetList(ctx context.Context, in *v1.JobLevelInfo, page, size int32) (*v1.GetListJobLevelRes, error) {
	res := &v1.GetListJobLevelRes{}
	resData := make([]*v1.JobLevelInfo, 0)
	jobLevelEntity := make([]entity.JobLevel, 0)

	query := dao.JobLevel.Ctx(ctx)

	if len(in.GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.JobLevel.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetName(), "%")})
	}
	if in.GetId() > 0 {
		query = query.Where(dao.JobLevel.Columns().Id, in.GetId())
	}

	if len(in.GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.JobLevel.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetRemark(), "%")})
	}

	query, totalSize, err := library.GetListWithPage(query, page, size)
	if err != nil {
		return res, err
	}
	err = query.Scan(&jobLevelEntity)
	jonLevelEntityByte, _ := json.Marshal(jobLevelEntity)
	json.Unmarshal(jonLevelEntityByte, &resData)

	res.Page = page
	res.Size = size
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}
func (s *sJobLevel) Modify(ctx context.Context, in *v1.JobLevelInfo) (*v1.JobLevelInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("职级名称不能为空")
	}
	if g.IsEmpty(in.GetId()) {
		return in, errors.New("请选择编辑的数据对象")
	}

	// 不能重名
	if !s.isUniqueJobLevel(ctx, in.GetName(), in.GetId()) {
		return in, errors.New("输入的职级信息重复，请确认信息是否正确")
	}

	data := do.JobLevel{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return in, err
	}

	data.UpdateTime = gtime.Now()
	if _, err = dao.JobLevel.Ctx(ctx).Where(dao.JobLevel.Columns().Id, in.GetId()).Data(data).Update(); err != nil {
		return in, err
	}
	return in, nil
}
func (s *sJobLevel) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.JobLevelInfo{Id: id})
	if (err != nil && err == sql.ErrNoRows) || g.IsNil(info) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	// 删除部门时，该部门下不能存在员工信息
	employeeInfo, err := service.Employee().GetOne(ctx, &v1.GetOneEmployeeReq{JobLevel: []int32{id}})
	if err != nil && err != sql.ErrNoRows {
		return false, "当前数据有误，请联系相关维护人员", err
	}
	if !g.IsNil(employeeInfo) {
		return false, "请先移除当前职级下的员工信息", errors.New(fmt.Sprintf("当前职级存在员工信息ID：%d,工号:%s", employeeInfo.Employee.Id, employeeInfo.Employee.WorkNumber))
	}

	_, err = dao.JobLevel.Ctx(ctx).Where(dao.JobLevel.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除部门职级失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sJobLevel) isUniqueJobLevel(ctx context.Context, name string, id int32) bool {
	var jobLevel *v1.JobLevelInfo
	err := dao.JobLevel.Ctx(ctx).
		Where(dao.JobLevel.Columns().Name, name).
		WhereNot(dao.JobLevel.Columns().Id, id).Scan(&jobLevel)
	if (err != nil && err != sql.ErrNoRows) || !g.IsNil(jobLevel) {
		return false
	}

	return true
}
