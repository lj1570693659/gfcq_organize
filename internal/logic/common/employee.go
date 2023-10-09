package common

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogo/protobuf/sortkeys"
	"github.com/lj1570693659/gfcq_product/internal/dao"
	"github.com/lj1570693659/gfcq_product/internal/library"
	"github.com/lj1570693659/gfcq_product/internal/model/do"
	"github.com/lj1570693659/gfcq_product/internal/model/entity"
	"github.com/lj1570693659/gfcq_product/internal/service"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"strings"
)

type (
	sEmployee struct{}
)

func init() {
	service.RegisterEmployee(&sEmployee{})
}

func (s *sEmployee) Create(ctx context.Context, in *v1.CreateEmployeeReq) (*v1.CreateEmployeeRes, error) {
	res := &v1.CreateEmployeeRes{Employee: &v1.EmployeeInfo{}}
	in, err := s.checkInputData(ctx, in)
	if err != nil {
		return res, err
	}

	// 工号不能重复
	_, isUnique := s.isUniqueWorkNumber(ctx, in.GetWorkNumber(), 0)
	if !isUnique {
		return res, errors.New("输入的工号重复，请确认信息是否正确")
	}

	data := do.Employee{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	// 根据岗位获取部门信息
	departs, err := s.getDepartsByJobs(ctx, in.JobId)
	if err != nil {
		return res, err
	}
	data.JobId = strings.Join(gconv.Strings(in.JobId), ",")
	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()

	// 增加员工信息
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		lastInsertId, err := dao.Employee.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			return err
		}
		res.Employee.Id = gconv.Int32(lastInsertId)
		// 创建员工-岗位关联信息
		for _, job := range departs {
			if _, err = service.EmployeeJob().Create(ctx, &v1.EmployeeJobInfo{
				EmployeeId: res.Employee.Id,
				JobId:      job.Id,
				DepartId:   job.DepartId,
			}); err != nil {
				return err
			}
		}
		return nil
	})

	return res, err
}

func (s *sEmployee) GetOne(ctx context.Context, in *v1.GetOneEmployeeReq) (*v1.GetOneEmployeeRes, error) {
	employee := &v1.EmployeeInfo{}
	res := &v1.GetOneEmployeeRes{}
	query := dao.Employee.Ctx(ctx)
	if in.GetId() > 0 {
		query = query.Where(dao.Employee.Columns().Id, in.GetId())
	}

	if len(in.GetUserName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().UserName), g.Slice{fmt.Sprintf("%s%s", in.GetUserName(), "%")})
	}

	if len(in.GetWorkNumber()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().WorkNumber), g.Slice{fmt.Sprintf("%s%s", in.GetWorkNumber(), "%")})
	}

	if len(in.GetEmail()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().Email), g.Slice{fmt.Sprintf("%s%s", in.GetEmail(), "%")})
	}

	if len(in.GetPhone()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().Phone), g.Slice{fmt.Sprintf("%s%s", in.GetPhone(), "%")})
	}

	if len(library.DeleteIntSlice(in.GetJobLevel())) > 0 {
		query = query.WhereIn(dao.Employee.Columns().JobLevel, in.GetJobLevel())
	}

	if in.GetStatus() > 0 {
		query = query.Where(dao.Employee.Columns().Status, in.GetStatus())
	}
	if len(in.DepartId) > 0 {
		queryBuilder := query.Builder().Where(fmt.Sprintf("FIND_IN_SET(%d,%s)", in.DepartId[0], dao.Employee.Columns().DepartId))
		for _, departId := range in.DepartId[1:] {
			queryBuilder = queryBuilder.WhereOr(fmt.Sprintf("FIND_IN_SET(%d,%s)", departId, dao.Employee.Columns().DepartId))
		}
		query = query.Where(queryBuilder)
	}

	err := query.Scan(&employee)
	res.Employee = employee
	if err != nil && err != sql.ErrNoRows {
		res.DepartString = employee.DepartId
		res.DepartIds = gconv.Int32s(strings.Split(employee.DepartId, ","))
		res.JobIds = gconv.Int32s(strings.Split(employee.JobId, ","))
		res.JobIdString = employee.JobId
	}
	return res, err
}

func (s *sEmployee) GetList(ctx context.Context, in *v1.EmployeeInfo, page, size int32) (*v1.GetListEmployeeRes, error) {
	res := &v1.GetListEmployeeRes{}
	resData := make([]*v1.EmployeeInfo, 0)
	employeeEntity := make([]entity.Employee, 0)

	query := dao.Employee.Ctx(ctx)

	if len(in.GetUserName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().UserName), g.Slice{fmt.Sprintf("%s%s", in.GetUserName(), "%")})
	}

	if len(in.GetWorkNumber()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().WorkNumber), g.Slice{fmt.Sprintf("%s%s", in.GetWorkNumber(), "%")})
	}

	if len(in.GetEmail()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().Email), g.Slice{fmt.Sprintf("%s%s", in.GetEmail(), "%")})
	}

	if len(in.GetPhone()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().Phone), g.Slice{fmt.Sprintf("%s%s", in.GetPhone(), "%")})
	}

	if len(in.DepartId) > 0 {
		departIds := library.DeleteIntSlice(gconv.Int32s(strings.Split(in.DepartId, ",")))
		if len(departIds) > 0 {
			queryBuilder := query.Builder().Where(fmt.Sprintf("FIND_IN_SET(%d,%s)", departIds[0], dao.Employee.Columns().DepartId))
			for _, departId := range departIds[1:] {
				queryBuilder = queryBuilder.WhereOr(fmt.Sprintf("FIND_IN_SET(%d,%s)", departId, dao.Employee.Columns().DepartId))
			}
			query = query.Where(queryBuilder)
		}

	}

	if len(in.JobId) > 0 {
		jobIds := library.DeleteIntSlice(gconv.Int32s(strings.Split(in.JobId, ",")))
		if len(jobIds) > 0 {
			queryBuilder := query.Builder().Where(fmt.Sprintf("FIND_IN_SET(%d,%s)", jobIds[0], dao.Employee.Columns().JobId))
			for _, jobId := range jobIds[1:] {
				queryBuilder = queryBuilder.WhereOr(fmt.Sprintf("FIND_IN_SET(%d,%s)", jobId, dao.Employee.Columns().JobId))
			}
			query = query.Where(queryBuilder)
		}
	}

	if in.GetJobLevel() > 0 {
		query = query.Where(dao.Employee.Columns().JobLevel, in.GetJobLevel())
	}
	if in.GetStatus() > 0 {
		query = query.Where(dao.Employee.Columns().Status, in.GetStatus())
	}

	query, totalSize, err := library.GetListWithPage(query, page, size)
	if err != nil {
		return res, err
	}
	err = query.OrderDesc(dao.Employee.Columns().Id).Scan(&employeeEntity)
	employeeEntityByte, _ := json.Marshal(employeeEntity)
	json.Unmarshal(employeeEntityByte, &resData)

	res.Page = page
	res.Size = size
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sEmployee) GetAll(ctx context.Context, in *v1.EmployeeInfo) (*v1.GetAllEmployeeRes, error) {
	res := &v1.GetAllEmployeeRes{}
	resData := make([]*v1.EmployeeInfo, 0)
	employeeEntity := make([]entity.Employee, 0)

	query := dao.Employee.Ctx(ctx)

	if len(in.GetUserName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().UserName), g.Slice{fmt.Sprintf("%s%s", in.GetUserName(), "%")})
	}

	if len(in.GetWorkNumber()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().WorkNumber), g.Slice{fmt.Sprintf("%s%s", in.GetWorkNumber(), "%")})
	}

	if len(in.GetEmail()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().Email), g.Slice{fmt.Sprintf("%s%s", in.GetEmail(), "%")})
	}

	if len(in.GetPhone()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Employee.Columns().Phone), g.Slice{fmt.Sprintf("%s%s", in.GetPhone(), "%")})
	}

	if len(in.DepartId) > 0 {
		departIds := library.DeleteIntSlice(gconv.Int32s(strings.Split(in.DepartId, ",")))
		if len(departIds) > 0 {
			queryBuilder := query.Builder().Where(fmt.Sprintf("FIND_IN_SET(%d,%s)", departIds[0], dao.Employee.Columns().DepartId))
			for _, departId := range departIds[1:] {
				queryBuilder = queryBuilder.WhereOr(fmt.Sprintf("FIND_IN_SET(%d,%s)", departId, dao.Employee.Columns().DepartId))
			}
			query = query.Where(queryBuilder)
		}

	}

	if len(in.JobId) > 0 {
		jobIds := library.DeleteIntSlice(gconv.Int32s(strings.Split(in.JobId, ",")))
		if len(jobIds) > 0 {
			queryBuilder := query.Builder().Where(fmt.Sprintf("FIND_IN_SET(%d,%s)", jobIds[0], dao.Employee.Columns().JobId))
			for _, jobId := range jobIds[1:] {
				queryBuilder = queryBuilder.WhereOr(fmt.Sprintf("FIND_IN_SET(%d,%s)", jobId, dao.Employee.Columns().JobId))
			}
			query = query.Where(queryBuilder)
		}
	}

	if in.GetJobLevel() > 0 {
		query = query.Where(dao.Employee.Columns().JobLevel, in.GetJobLevel())
	}
	if in.GetStatus() > 0 {
		query = query.Where(dao.Employee.Columns().Status, in.GetStatus())
	}

	err := query.Scan(&employeeEntity)
	employeeEntityByte, _ := json.Marshal(employeeEntity)
	json.Unmarshal(employeeEntityByte, &resData)

	res.Data = resData
	return res, err
}

func (s *sEmployee) Modify(ctx context.Context, in *v1.ModifyEmployeeReq) (*v1.ModifyEmployeeRes, error) {
	res := &v1.ModifyEmployeeRes{}
	if in.GetId() == 0 {
		return res, errors.New("请选择编辑的数据对象")
	}
	checkIn, err := s.checkInputData(ctx, &v1.CreateEmployeeReq{
		UserName:   in.GetUserName(),
		JobId:      in.GetJobId(),
		WorkNumber: in.GetWorkNumber(),
		Status:     in.GetStatus(),
		Sex:        in.GetSex(),
		JobLevel:   in.GetJobLevel(),
	})
	if err != nil {
		return res, err
	}

	// 工号不能重复
	_, isUnique := s.isUniqueWorkNumber(ctx, in.GetWorkNumber(), in.GetId())
	if !isUnique {
		return res, errors.New("输入的工号重复，请确认信息是否正确")
	}

	data := do.Employee{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	// 判断岗位和部门信息是否调整
	employee, err := s.GetOne(ctx, &v1.GetOneEmployeeReq{
		WorkNumber: in.GetWorkNumber(),
	})
	if err != nil {
		return res, err
	}

	// 员工岗位变动
	data.JobId = strings.Join(gconv.Strings(checkIn.JobId), ",")
	data.DepartId = employee.Employee.DepartId
	departIds := make([]int32, 0)
	if data.JobId != employee.JobIdString {
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 根据岗位获取部门信息
			departs, err := s.getDepartsByJobs(ctx, in.JobId)
			if err != nil {
				return err
			}

			// 删除原有员工-岗位关联信息
			isDeleted, msg, err := service.EmployeeJob().Delete(ctx, &v1.DeleteEmployeeJobReq{
				EmployeeId: in.Id,
			})
			if err != nil || !isDeleted {
				return errors.New(msg)
			}
			// 创建员工-岗位关联信息
			for _, job := range departs {
				if _, err = service.EmployeeJob().Create(ctx, &v1.EmployeeJobInfo{
					EmployeeId: in.Id,
					JobId:      job.Id,
					DepartId:   job.DepartId,
				}); err != nil {
					return err
				}
				departIds = append(departIds, job.DepartId)
			}
			return nil
		})
		if err != nil {
			return res, err
		}

		data.DepartId = strings.Join(gconv.Strings(departIds), ",")
	}

	data.UpdateTime = gtime.Now()
	if _, err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().Id, in.GetId()).Data(data).Update(); err != nil {
		return res, err
	}

	return res, nil
}

func (s *sEmployee) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.GetOneEmployeeReq{Id: id})
	if (err != nil && err == sql.ErrNoRows) || g.IsNil(info) {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	if info.Employee.Status != v1.StatusEnum_terminated {
		return false, "只能删除已离职员工信息", errors.New(fmt.Sprintf("当前员工未离职，员工信息工号：%s,当前状态:%s", info.Employee.WorkNumber, info.Employee.Status))
	}

	_, err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除员工数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}

func (s *sEmployee) checkInputData(ctx context.Context, in *v1.CreateEmployeeReq) (*v1.CreateEmployeeReq, error) {
	if in.GetJobLevel() == 0 {
		return in, errors.New("请输入正确格式的职级信息")
	}
	if len(in.GetJobId()) == 0 {
		return in, errors.New("请选择对应岗位")
	}
	if len(in.GetWorkNumber()) == 0 {
		return in, errors.New("请输入员工工号")
	}
	if len(in.GetUserName()) == 0 {
		return in, errors.New("请输入员工姓名")
	}
	if in.GetStatus() > v1.StatusEnum_terminated || in.GetStatus() < v1.StatusEnum_anything {
		return in, errors.New("请选择正确的员工状态")
	}
	if in.GetSex() > v1.SexEnum_woman || in.GetSex() < v1.SexEnum_unknow {
		return in, errors.New("请选择正确的员工性别")
	}

	// 职级数据是否正确校验
	jobLevelInfo, err := service.JobLevel().GetOne(ctx, &v1.JobLevelInfo{Id: in.GetJobLevel()})
	if (err != nil && err == sql.ErrNoRows) || g.IsNil(jobLevelInfo) {
		return in, errors.New("选择的职级信息不存在，请再次确认")
	}

	// 校验岗位信息是否正确
	departIds := make([]string, 0)
	for _, jobId := range in.GetJobId() {
		jobInfo, err := service.Job().GetOne(ctx, &v1.JobInfo{Id: jobId})
		if (err != nil && err.Error() == sql.ErrNoRows.Error()) || g.IsNil(jobInfo) {
			return in, errors.New("选择的岗位信息不存在，请再次确认")
		}
		departIds = append(departIds, gconv.String(jobInfo.DepartId))
	}

	in.DepartId = strings.Join(departIds, ",")
	sortkeys.Int32s(in.JobId)

	return in, nil
}

func (s *sEmployee) isUniqueWorkNumber(ctx context.Context, workNumber string, id int32) (*v1.EmployeeInfo, bool) {
	// 工号不能重复
	var employee *v1.EmployeeInfo
	query := dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().WorkNumber, workNumber)
	if !g.IsEmpty(id) {
		query = query.WhereNot(dao.Employee.Columns().Id, id)
	}
	err := query.Scan(&employee)
	if (err != nil && err.Error() != sql.ErrNoRows.Error()) || !g.IsNil(employee) {
		return employee, false
	}

	return employee, true
}

// 根据岗位获取部门信息
func (s *sEmployee) getDepartsByJobs(ctx context.Context, jobIds []int32) ([]*v1.JobInfo, error) {
	// 工号不能重复
	departs := make([]*v1.JobInfo, 0)
	if len(jobIds) == 0 {
		return departs, nil
	}
	for _, jobId := range jobIds {
		jobInfo, err := service.Job().GetOne(ctx, &v1.JobInfo{
			Id: jobId,
		})

		if err != nil {
			return departs, err
		}
		departs = append(departs, jobInfo)
	}

	return departs, nil
}
