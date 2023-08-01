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
	sJob struct{}
)

func init() {
	service.RegisterJob(&sJob{})
}

func (s *sJob) Create(ctx context.Context, in *v1.JobInfo) (*v1.JobInfo, error) {
	in, err := s.checkInputData(ctx, in)
	if err != nil {
		return in, err
	}

	// 同部门下岗位不能重复
	job, err := s.GetOne(ctx, &v1.JobInfo{DepartId: in.GetDepartId(), Name: in.GetName()})
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if !g.IsNil(job) && job.Id > 0 {
		return in, errors.New("当前部门下岗位名称已存在，请确认信息是否正确")
	}

	data := do.Job{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return in, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.Job.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return in, err
	}
	in.Id = gconv.Int32(lastInsertId)
	return in, nil
}
func (s *sJob) GetOne(ctx context.Context, in *v1.JobInfo) (*v1.JobInfo, error) {
	var job *v1.JobInfo
	query := dao.Job.Ctx(ctx)

	if len(in.GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Job.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetName(), "%")})
	}
	if in.GetId() > 0 {
		query = query.Where(dao.Job.Columns().Id, in.GetId())
	}
	if in.GetDepartId() > 0 {
		query = query.Where(dao.Job.Columns().DepartId, in.GetDepartId())
	}

	err := query.Scan(&job)

	return job, err
}

func (s *sJob) GetList(ctx context.Context, in *v1.JobInfo, page, size int32) (*v1.GetListJobRes, error) {
	res := &v1.GetListJobRes{}
	resData := make([]*v1.JobInfo, 0)
	jobEntity := make([]entity.Job, 0)

	query := dao.Job.Ctx(ctx)

	if len(in.GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Job.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetName(), "%")})
	}

	if len(in.GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Job.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetRemark(), "%")})
	}

	if in.GetDepartId() > 0 {
		query = query.Where(dao.Employee.Columns().DepartId, in.GetDepartId())
	}

	query, totalSize, err := library.GetListWithPage(query, page, size)
	if err != nil {
		return res, err
	}
	err = query.Scan(&jobEntity)
	jobEntityByte, _ := json.Marshal(jobEntity)
	json.Unmarshal(jobEntityByte, &resData)

	res.Page = page
	res.Size = size
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sJob) GetAll(ctx context.Context, in *v1.JobInfo) (*v1.GetAllJobRes, error) {
	res := &v1.GetAllJobRes{}
	resData := make([]*v1.JobInfo, 0)
	jobEntity := make([]entity.Job, 0)

	query := dao.Job.Ctx(ctx)

	if len(in.GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Job.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetName(), "%")})
	}

	if len(in.GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Job.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetRemark(), "%")})
	}

	if in.GetDepartId() > 0 {
		query = query.Where(dao.Employee.Columns().DepartId, in.GetDepartId())
	}

	err := query.Scan(&jobEntity)
	jobEntityByte, _ := json.Marshal(jobEntity)
	json.Unmarshal(jobEntityByte, &resData)

	res.Data = resData
	return res, err
}

func (s *sJob) Modify(ctx context.Context, in *v1.JobInfo) (*v1.JobInfo, error) {
	if in.GetId() == 0 {
		return in, errors.New("请选择编辑的数据对象")
	}

	in, err := s.checkInputData(ctx, in)
	if err != nil {
		return in, err
	}

	// 同部门下岗位不能重复
	if !s.isUniqueName(ctx, in.GetName(), in.GetDepartId(), in.GetId()) {
		return in, errors.New("当前部门下岗位名称已存在，请确认信息是否正确")
	}

	data := do.Job{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return in, err
	}

	data.UpdateTime = gtime.Now()
	if _, err = dao.Job.Ctx(ctx).Where(dao.Job.Columns().Id, in.GetId()).Data(data).Update(); err != nil {
		return in, err
	}
	return in, nil
}
func (s *sJob) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.JobInfo{Id: id})
	if (err != nil && err == sql.ErrNoRows) || g.IsNil(info) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	// 删除岗位时，该岗位下不能存在员工信息 TODO
	employeeInfo, err := service.EmployeeJob().GetOne(ctx, &v1.EmployeeJobInfo{JobId: id})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return false, "当前数据有误，请联系相关维护人员", err
	}
	if !g.IsNil(employeeInfo) {
		return false, "请先移除当前岗位下的员工信息", errors.New(fmt.Sprintf("当前岗位存在员工信息ID：%d,员工信息:%d", employeeInfo.Id, employeeInfo.EmployeeId))
	}

	_, err = dao.Job.Ctx(ctx).Where(dao.Job.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除岗位数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sJob) checkInputData(ctx context.Context, in *v1.JobInfo) (*v1.JobInfo, error) {
	if in.GetDepartId() == 0 {
		return in, errors.New("请选择所在部门")
	}
	if len(in.GetName()) == 0 {
		return in, errors.New("请输入岗位名称")
	}

	// 数据是否正确校验
	departmentInfo, err := service.Department().GetOne(ctx, &v1.DepartmentInfo{Id: in.GetDepartId()})
	if (err != nil && err == sql.ErrNoRows) || g.IsNil(departmentInfo) {
		return in, errors.New("选择的部门信息不存在，请再次确认")
	}

	return in, nil
}

func (s *sJob) isUniqueName(ctx context.Context, name string, departId, id int32) bool {
	// 工号不能重复
	var job *v1.JobInfo
	err := dao.Job.Ctx(ctx).
		Where(dao.Job.Columns().DepartId, departId).
		Where(dao.Job.Columns().Name, name).
		WhereNot(dao.Job.Columns().Id, id).Scan(&job)
	if (err != nil && err != sql.ErrNoRows) || !g.IsNil(job) {
		return false
	}

	return true
}
