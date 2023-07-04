// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// JobDao is the data access object for table cqgf_job.
type JobDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns JobColumns // columns contains all the column names of Table for convenient usage.
}

// JobColumns defines and stores column names for table cqgf_job.
type JobColumns struct {
	Id         string //
	Name       string // 岗位名称
	DepartId   string // 所属部门
	Remark     string // 预留备注信息
	CreateTime string // 数据新增时间
	UpdateTime string // 最后一次更新数据时间
}

// jobColumns holds the columns for table cqgf_job.
var jobColumns = JobColumns{
	Id:         "id",
	Name:       "name",
	DepartId:   "depart_id",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewJobDao creates and returns a new DAO object for table data access.
func NewJobDao() *JobDao {
	return &JobDao{
		group:   "default",
		table:   "cqgf_job",
		columns: jobColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *JobDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *JobDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *JobDao) Columns() JobColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *JobDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *JobDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *JobDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
