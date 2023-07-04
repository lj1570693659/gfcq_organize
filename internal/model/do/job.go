// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Job is the golang structure of table cqgf_job for DAO operations like Where/Data.
type Job struct {
	g.Meta     `orm:"table:cqgf_job, do:true"`
	Id         interface{} //
	Name       interface{} // 岗位名称
	DepartId   interface{} // 所属部门
	Remark     interface{} // 预留备注信息
	CreateTime *gtime.Time // 数据新增时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
