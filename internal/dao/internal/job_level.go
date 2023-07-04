// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// JobLevelDao is the data access object for table cqgf_job_level.
type JobLevelDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns JobLevelColumns // columns contains all the column names of Table for convenient usage.
}

// JobLevelColumns defines and stores column names for table cqgf_job_level.
type JobLevelColumns struct {
	Id         string //
	Name       string // 职级名称
	Remark     string // 预留备注信息
	CreateTime string // 数据新增时间
	UpdateTime string // 最后一次更新数据时间
}

// jobLevelColumns holds the columns for table cqgf_job_level.
var jobLevelColumns = JobLevelColumns{
	Id:         "id",
	Name:       "name",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewJobLevelDao creates and returns a new DAO object for table data access.
func NewJobLevelDao() *JobLevelDao {
	return &JobLevelDao{
		group:   "default",
		table:   "cqgf_job_level",
		columns: jobLevelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *JobLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *JobLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *JobLevelDao) Columns() JobLevelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *JobLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *JobLevelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *JobLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
