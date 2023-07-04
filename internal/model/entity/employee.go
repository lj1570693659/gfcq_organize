// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Employee is the golang structure for table employee.
type Employee struct {
	Id           int         `json:"id"           description:""`
	UserName     string      `json:"userName"     description:"员工姓名"`
	WorkNumber   string      `json:"workNumber"   description:"员工工号"`
	Sex          uint        `json:"sex"          description:"性别（0：未知 1：男  2：女）"`
	Phone        string      `json:"phone"        description:"手机号码"`
	Email        string      `json:"email"        description:"邮箱号码"`
	DepartId     string      `json:"departId"     description:"所属部门"`
	JobLevel     uint        `json:"jobLevel"     description:"职级"`
	JobId        string      `json:"jobId"        description:"岗位信息"`
	InstructorId int         `json:"instructorId" description:"指导老师"`
	Status       int         `json:"status"       description:"在职状态（1：在职 2：试用期 3：实习期 4：已离职）"`
	Remark       string      `json:"remark"       description:"预留备注信息"`
	CreateTime   *gtime.Time `json:"createTime"   description:"新增数据时间"`
	UpdateTime   *gtime.Time `json:"updateTime"   description:"最后一次更新数据时间"`
}
