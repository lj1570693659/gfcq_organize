package common

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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
	sEmployeeJob struct{}
)

func init() {
	service.RegisterEmployeeJob(&sEmployeeJob{})
}

func (s *sEmployeeJob) Create(ctx context.Context, in *v1.EmployeeJobInfo) (*v1.EmployeeJobInfo, error) {
	in, err := s.checkInputData(ctx, in)
	if err != nil {
		return in, err
	}

	// 员工&岗位组合唯一
	employeeJob, err := s.GetOne(ctx, &v1.EmployeeJobInfo{EmployeeId: in.GetEmployeeId(), JobId: in.GetJobId()})
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if !g.IsNil(employeeJob) && employeeJob.Id > 0 {
		return in, errors.New("岗前员工对应岗位信息已存在，请确认信息是否正确")
	}

	data := do.EmployeeJob{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return in, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := dao.EmployeeJob.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return in, err
	}
	in.Id = gconv.Int32(lastInsertId)
	return in, nil
}
func (s *sEmployeeJob) GetOne(ctx context.Context, in *v1.EmployeeJobInfo) (*v1.EmployeeJobInfo, error) {
	var employeeJob *v1.EmployeeJobInfo
	query := dao.EmployeeJob.Ctx(ctx)

	if in.GetEmployeeId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().EmployeeId, in.GetEmployeeId())
	}
	if in.GetJobId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().JobId, in.GetJobId())
	}
	if in.GetId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().Id, in.GetId())
	}

	err := query.Scan(&employeeJob)

	return employeeJob, err
}
func (s *sEmployeeJob) GetList(ctx context.Context, in *v1.EmployeeJobInfo, page, size int32) (*v1.GetListEmployeeJobRes, error) {
	res := &v1.GetListEmployeeJobRes{}
	resData := make([]*v1.EmployeeJobInfo, 0)
	employeeJobEntity := make([]entity.EmployeeJob, 0)

	query := dao.EmployeeJob.Ctx(ctx)

	if in.GetEmployeeId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().EmployeeId, in.GetEmployeeId())
	}
	if in.GetJobId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().JobId, in.GetJobId())
	}
	if in.GetId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().Id, in.GetId())
	}

	query, totalSize, err := library.GetListWithPage(query, page, size)
	if err != nil {
		return res, err
	}
	err = query.Scan(&employeeJobEntity)
	employeeJobEntityByte, _ := json.Marshal(employeeJobEntity)
	json.Unmarshal(employeeJobEntityByte, &resData)

	res.Page = page
	res.Size = size
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sEmployeeJob) GetCount(ctx context.Context, in *v1.GetCountEmployeeJobReq) (*v1.GetCountEmployeeJobRes, error) {
	res := &v1.GetCountEmployeeJobRes{}
	query := dao.EmployeeJob.Ctx(ctx)

	if in.GetEmployeeJob().GetEmployeeId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().EmployeeId, in.GetEmployeeJob().GetEmployeeId())
	}
	if in.GetEmployeeJob().GetDepartId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().DepartId, in.GetEmployeeJob().GetDepartId())
	}
	if in.GetEmployeeJob().GetJobId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().JobId, in.GetEmployeeJob().GetJobId())
	}
	if in.GetEmployeeJob().GetId() > 0 {
		query = query.Where(dao.EmployeeJob.Columns().Id, in.GetEmployeeJob().GetId())
	}
	count, err := query.Group(in.GetGroupBy()).Count(in.GetGetFiledNameCount())
	res.Count = gconv.Int32(count)
	return res, err
}

func (s *sEmployeeJob) Modify(ctx context.Context, in *v1.EmployeeJobInfo) (*v1.EmployeeJobInfo, error) {
	if g.IsEmpty(in.GetId()) {
		return in, errors.New("当前操作的数据有误，请联系相关维护人员")
	}

	in, err := s.checkInputData(ctx, in)
	if err != nil {
		return in, err
	}

	// 员工&岗位组合唯一
	var employeeJob *v1.EmployeeJobInfo
	err = dao.EmployeeJob.Ctx(ctx).
		Where(dao.EmployeeJob.Columns().EmployeeId, in.GetEmployeeId()).
		Where(dao.EmployeeJob.Columns().JobId, in.GetJobId()).
		WhereNot(dao.EmployeeJob.Columns().Id, in.GetId()).Scan(&employeeJob)
	if (err != nil && err != sql.ErrNoRows) || !g.IsNil(employeeJob) {
		return in, errors.New("岗前员工对应岗位信息已存在，请确认信息是否正确")
	}

	data := do.EmployeeJob{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return in, err
	}

	data.UpdateTime = gtime.Now()
	if _, err = dao.EmployeeJob.Ctx(ctx).Where(dao.EmployeeJob.Columns().Id, in.GetId()).Data(data).Update(); err != nil {
		return in, err
	}

	return in, nil
}
func (s *sEmployeeJob) Delete(ctx context.Context, info *v1.DeleteEmployeeJobReq) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(info.Id) && g.IsEmpty(info.EmployeeId) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的数据为空")
	}

	query := dao.EmployeeJob.Ctx(ctx)
	// 校验修改的原始数据是否存在
	if !g.IsEmpty(info.Id) {
		employeeInfo, err := s.GetOne(ctx, &v1.EmployeeJobInfo{Id: info.Id})
		if (err != nil && err == sql.ErrNoRows) || g.IsNil(employeeInfo) {
			return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
		}
		query = query.Where(dao.EmployeeJob.Columns().Id, info.Id)
	} else if !g.IsEmpty(info.EmployeeId) {
		employeeInfo, err := s.GetOne(ctx, &v1.EmployeeJobInfo{EmployeeId: info.EmployeeId})
		if (err != nil && err.Error() != sql.ErrNoRows.Error()) || g.IsNil(employeeInfo) {
			g.Log("employee").Info(ctx, "接收到的EmployeeId在数据库中没有对应数据")
		}
		query = query.Where(dao.EmployeeJob.Columns().EmployeeId, info.EmployeeId)
	}

	if _, err = query.Delete(); err != nil {
		return false, "删除员工&岗位数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sEmployeeJob) checkInputData(ctx context.Context, in *v1.EmployeeJobInfo) (*v1.EmployeeJobInfo, error) {
	if in.GetEmployeeId() == 0 {
		return in, errors.New("请选择员工信息")
	}
	if in.GetJobId() == 0 {
		return in, errors.New("请选择员工对应岗位信息")
	}

	// 员工数据是否正确校验
	employeeInfo, err := service.Employee().GetOne(ctx, &v1.GetOneEmployeeReq{Id: in.GetEmployeeId()})
	if (err != nil && err == sql.ErrNoRows) || g.IsNil(employeeInfo) {
		return in, errors.New("选择的员工信息不存在，请再次确认")
	}

	// 校验岗位信息是否正确
	jobInfo, err := service.Job().GetOne(ctx, &v1.JobInfo{Id: in.GetJobId()})
	if (err != nil && err == sql.ErrNoRows) || g.IsNil(jobInfo) {
		return in, errors.New("选择的岗位信息不存在，请再次确认")
	}

	return in, nil
}
