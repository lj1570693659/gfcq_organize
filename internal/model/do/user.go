// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   	`orm:"table:cqgf_user, do:true"`
	Id       	interface{} // User ID
	Passport 	interface{} // User Passport
	Password 	interface{} // User Password
	Nickname 	interface{} // User Nickname
	UserName 	interface{} // User Nickname
	EmployeeId 	interface{} // User Nickname
	WorkNumber interface{} // User Nickname
	CreateTime interface{} // Created Time
	UpdateTime interface{} // Updated Time
}
