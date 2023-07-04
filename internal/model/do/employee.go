// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Employee is the golang structure of table cqgf_employee for DAO operations like Where/Data.
type Employee struct {
	g.Meta       `orm:"table:cqgf_employee, do:true"`
	Id           interface{} //
	UserName     interface{} // 员工姓名
	WorkNumber   interface{} // 员工工号
	Sex          interface{} // 性别（0：未知 1：男  2：女）
	Phone        interface{} // 手机号码
	Email        interface{} // 邮箱号码
	DepartId     interface{} // 所属部门
	JobLevel     interface{} // 职级
	JobId        interface{} // 岗位信息
	InstructorId interface{} // 指导老师
	Status       interface{} // 在职状态（1：在职 2：试用期 3：实习期 4：已离职）
	Remark       interface{} // 预留备注信息
	CreateTime   *gtime.Time // 新增数据时间
	UpdateTime   *gtime.Time // 最后一次更新数据时间
}
