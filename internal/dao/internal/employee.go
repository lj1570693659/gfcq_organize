// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeDao is the data access object for table cqgf_employee.
type EmployeeDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns EmployeeColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeeColumns defines and stores column names for table cqgf_employee.
type EmployeeColumns struct {
	Id           string //
	UserName     string // 员工姓名
	WorkNumber   string // 员工工号
	Sex          string // 性别（0：未知 1：男  2：女）
	Phone        string // 手机号码
	Email        string // 邮箱号码
	DepartId     string // 所属部门
	JobLevel     string // 职级
	JobId        string // 岗位信息
	InstructorId string // 指导老师
	Status       string // 在职状态（1：在职 2：试用期 3：实习期 4：已离职）
	Remark       string // 预留备注信息
	CreateTime   string // 新增数据时间
	UpdateTime   string // 最后一次更新数据时间
}

// employeeColumns holds the columns for table cqgf_employee.
var employeeColumns = EmployeeColumns{
	Id:           "id",
	UserName:     "user_name",
	WorkNumber:   "work_number",
	Sex:          "sex",
	Phone:        "phone",
	Email:        "email",
	DepartId:     "depart_id",
	JobLevel:     "job_level",
	JobId:        "job_id",
	InstructorId: "instructor_id",
	Status:       "status",
	Remark:       "remark",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewEmployeeDao creates and returns a new DAO object for table data access.
func NewEmployeeDao() *EmployeeDao {
	return &EmployeeDao{
		group:   "default",
		table:   "cqgf_employee",
		columns: employeeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EmployeeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EmployeeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EmployeeDao) Columns() EmployeeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EmployeeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EmployeeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EmployeeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
