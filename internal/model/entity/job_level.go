// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// JobLevel is the golang structure for table job_level.
type JobLevel struct {
	Id         int         `json:"id"         description:""`
	Name       string      `json:"name"       description:"职级名称"`
	Remark     string      `json:"remark"     description:"预留备注信息"`
	CreateTime *gtime.Time `json:"createTime" description:"数据新增时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新数据时间"`
}
