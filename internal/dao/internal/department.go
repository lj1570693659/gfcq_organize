// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DepartmentDao is the data access object for table cqgf_department.
type DepartmentDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns DepartmentColumns // columns contains all the column names of Table for convenient usage.
}

// DepartmentColumns defines and stores column names for table cqgf_department.
type DepartmentColumns struct {
	Id         string //
	Name       string // 部门名称
	Pid        string // 上级部门
	Remark     string // 预留备注信息
	CreateTime string // 数据新增时间
	UpdateTime string // 最后一次更新数据时间
}

// departmentColumns holds the columns for table cqgf_department.
var departmentColumns = DepartmentColumns{
	Id:         "id",
	Name:       "name",
	Pid:        "pid",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewDepartmentDao creates and returns a new DAO object for table data access.
func NewDepartmentDao() *DepartmentDao {
	return &DepartmentDao{
		group:   "default",
		table:   "cqgf_department",
		columns: departmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DepartmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DepartmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DepartmentDao) Columns() DepartmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DepartmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DepartmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DepartmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
