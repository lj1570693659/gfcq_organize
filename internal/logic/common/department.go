package common

import (
	"context"
	"database/sql"
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
	sDepartment struct{}
)

func init() {
	service.RegisterDepartment(&sDepartment{})
}

func (s *sDepartment) Create(ctx context.Context, in *v1.DepartmentInfo) (*v1.DepartmentInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("部门名称不能为空")
	}

	// 上级部门存在时，数据校验
	if in.GetPid() > 0 {
		//pidInfo, err := s.GetOne(ctx, &v1.DepartmentInfo{
		//	Id: in.GetPid(),
		//})
		//if (err != nil && err == sql.ErrNoRows) || pidInfo == nil {
		//	return in, errors.New("选择的上级部门不存在，请确认是否增加")
		//}
	}

	// 同级部门不能重名
	info, err := s.GetOne(ctx, &v1.DepartmentInfo{
		Pid:  in.GetPid(),
		Name: in.GetName(),
	})

	if (err != nil && err != sql.ErrNoRows) || !g.IsNil(info) {
		return in, err
	}
	if info != nil && info.Id > 0 {
		return in, errors.New("该部门下已存在同名部门，请重新命名")
	}
	data := do.Department{
		Id:               in.Id,
		Pid:              in.Pid,
		Name:             in.Name,
		NameEn:           in.NameEn,
		DepartmentLeader: in.DepartmentLeader,
		Remark:           in.Remark,
		Level:            in.Level,
		CreateTime:       gtime.Now(),
		UpdateTime:       gtime.Now(),
	}

	lastInsertId, err := dao.Department.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return in, err
	}
	in.Id = gconv.Int32(lastInsertId)
	return in, nil
}

func (s *sDepartment) GetOne(ctx context.Context, in *v1.DepartmentInfo) (*v1.DepartmentInfo, error) {
	fmt.Println("depart------------------GetOne--------------------", in.GetId())
	var depart *v1.DepartmentInfo
	query := dao.Department.Ctx(ctx)

	if len(in.GetName()) > 0 {
		//query = query.Where(fmt.Sprintf("%s like ?", dao.Department.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetName(), "%")})
		query = query.Where(dao.Department.Columns().Name, in.GetName())
	}
	if in.GetId() > 0 {
		query = query.Where(dao.Department.Columns().Id, in.GetId())
	}
	if in.GetPid() > 0 {
		query = query.Where(dao.Department.Columns().Pid, in.GetPid())
	}

	err := query.Scan(&depart)
	return depart, err
}

func (s *sDepartment) GetList(ctx context.Context, in *v1.DepartmentInfo, page, size int32) (*v1.GetListDepartmentRes, error) {
	res := &v1.GetListDepartmentRes{}
	resData := make([]*v1.DepartmentInfo, 0)
	depart := make([]entity.Department, 0)

	query := dao.Department.Ctx(ctx)

	if len(in.GetName()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Department.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.GetName(), "%")})
	}
	if in.GetPid() > 0 {
		query = query.Where(dao.Department.Columns().Pid, in.GetPid())
	}

	if len(in.GetRemark()) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", dao.Department.Columns().Remark), g.Slice{fmt.Sprintf("%s%s", in.GetRemark(), "%")})
	}

	query, totalSize, err := library.GetListWithPage(query, page, size)
	if err != nil {
		return res, err
	}
	err = query.Scan(&depart)
	for _, v := range depart {
		resData = append(resData, &v1.DepartmentInfo{
			Id:         gconv.Int32(v.Id),
			Name:       v.Name,
			Pid:        gconv.Int32(v.Pid),
			Remark:     v.Remark,
			CreateTime: v.CreateTime.String(),
			UpdateTime: v.UpdateTime.String(),
		})
	}
	res.Page = page
	res.Size = size
	res.TotalSize = totalSize
	res.Data = resData
	return res, err
}

func (s *sDepartment) GetListWithoutPage(ctx context.Context, in *v1.DepartmentInfo) (res *v1.GetListWithoutDepartmentRes, err error) {
	resData := make([]*v1.DepartmentInfo, 0)

	query := dao.Department.Ctx(ctx)
	condition := g.Map{}

	if len(in.GetName()) > 0 {
		condition[fmt.Sprintf("%s like ?", dao.Department.Columns().Name)] = fmt.Sprintf("%s%s%s", "%", in.GetName(), "%")
	}

	if in.GetId() > 0 {
		condition[dao.Department.Columns().Id] = in.GetId()
	}

	if in.GetPid() > 0 {
		condition[dao.Department.Columns().Pid] = in.GetPid()
	} else if in.GetPid() == -1 {
		condition[dao.Department.Columns().Pid] = 0
	}

	err = query.Where(condition).Scan(&resData)
	if err != nil {
		return res, err
	}

	return &v1.GetListWithoutDepartmentRes{
		Data: resData,
	}, err
}

func (s *sDepartment) Modify(ctx context.Context, in *v1.DepartmentInfo) (*v1.DepartmentInfo, error) {
	if len(in.GetName()) == 0 {
		return in, errors.New("部门名称不能为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.DepartmentInfo{
		Id: in.GetId(),
	})
	if (err != nil && err == sql.ErrNoRows) || info == nil {
		return in, errors.New("当前数据不存在，请确认是否增加")
	} else if err != nil && err != sql.ErrNoRows {
		return in, err
	}

	// 上级部门存在时，数据校验
	if in.GetPid() > 0 {
		pidInfo, err := s.GetOne(ctx, &v1.DepartmentInfo{
			Id: in.GetPid(),
		})
		if (err != nil && err == sql.ErrNoRows) || pidInfo == nil {
			return in, errors.New("选择的上级部门不存在，请确认是否增加")
		}
	}

	// 同级部门不能重名
	dName, err := s.GetOne(ctx, &v1.DepartmentInfo{
		Pid:  in.GetPid(),
		Name: in.GetName(),
	})
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if dName != nil && dName.Id != in.GetId() {
		return in, errors.New("该部门下已存在同名部门，请重新命名")
	}

	data := do.Department{
		Pid:              in.Pid,
		Name:             in.Name,
		NameEn:           in.NameEn,
		DepartmentLeader: in.DepartmentLeader,
		Level:            in.Level,
		Remark:           in.Remark,
		UpdateTime:       gtime.Now(),
	}

	_, err = dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, in.GetId()).Data(data).Update()
	if err != nil {
		return in, err
	}

	return in, nil
}

func (s *sDepartment) Delete(ctx context.Context, id int32) (isSuccess bool, msg string, err error) {
	if g.IsEmpty(id) {
		return false, "当前操作的数据有误，请联系相关维护人员", errors.New("接收到的ID数据为空")
	}

	// 校验修改的原始数据是否存在
	info, err := s.GetOne(ctx, &v1.DepartmentInfo{Id: id})
	if (err != nil && err == sql.ErrNoRows) || info == nil {
		return false, "当前数据不存在，请联系相关维护人员", errors.New("接收到的ID在数据库中没有对应数据")
	}

	// 删除父级部门时，校验子部门是否为空
	if g.IsEmpty(info.Pid) {
		zidInfo, err := s.GetOne(ctx, &v1.DepartmentInfo{Pid: id})
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return false, "当前数据不存在，请联系相关维护人员", err
		}

		if !g.IsNil(zidInfo) && !g.IsEmpty(zidInfo.Id) {
			return false, "请先移除当前部门下的子部门信息", errors.New(fmt.Sprintf("当前部门存在子部门信息ID：%d,name:%s", zidInfo.Id, zidInfo.Name))
		}
	}

	// 删除部门时，该部门下不能存在员工信息
	employeeInfo, err := service.Employee().GetOne(ctx, &v1.GetOneEmployeeReq{DepartId: []int32{id}})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return false, "当前数据有误，请联系相关维护人员", err
	}

	if !g.IsNil(employeeInfo) && !g.IsEmpty(employeeInfo.GetEmployee()) {
		return false, "请先移除当前部门下的员工信息", errors.New(fmt.Sprintf("当前部门存在员工信息ID：%d,工号:%s", employeeInfo.Employee.Id, employeeInfo.Employee.WorkNumber))
	}

	_, err = dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, id).Delete()
	if err != nil {
		return false, "删除部门数据失败，请联系相关维护人员", err
	}
	return true, "", nil
}
