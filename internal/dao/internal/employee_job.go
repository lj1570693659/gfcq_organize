// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeJobDao is the data access object for table cqgf_employee_job.
type EmployeeJobDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns EmployeeJobColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeeJobColumns defines and stores column names for table cqgf_employee_job.
type EmployeeJobColumns struct {
	Id         string //
	EmployeeId string // 员工信息
	JobId      string // 岗位信息
	DepartId   string // 所属部门
	Remark     string // 预留备注信息
	CreateTime string // 数据新增时间
	UpdateTime string // 最后一次更新数据时间
}

// employeeJobColumns holds the columns for table cqgf_employee_job.
var employeeJobColumns = EmployeeJobColumns{
	Id:         "id",
	EmployeeId: "employee_id",
	JobId:      "job_id",
	DepartId:   "depart_id",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewEmployeeJobDao creates and returns a new DAO object for table data access.
func NewEmployeeJobDao() *EmployeeJobDao {
	return &EmployeeJobDao{
		group:   "default",
		table:   "cqgf_employee_job",
		columns: employeeJobColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EmployeeJobDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EmployeeJobDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EmployeeJobDao) Columns() EmployeeJobColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EmployeeJobDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EmployeeJobDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EmployeeJobDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
