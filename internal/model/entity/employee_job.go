// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeJob is the golang structure for table employee_job.
type EmployeeJob struct {
	Id         int         `json:"id"         description:""`
	EmployeeId int         `json:"employeeId" description:"员工信息"`
	JobId      int         `json:"jobId"      description:"岗位信息"`
	DepartId   int         `json:"departId"   description:"所属部门"`
	Remark     string      `json:"remark"     description:"预留备注信息"`
	CreateTime *gtime.Time `json:"createTime" description:"数据新增时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新数据时间"`
}
