// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_product/internal/dao/internal"
)

// internalEmployeeDao is internal type for wrapping internal DAO implements.
type internalEmployeeDao = *internal.EmployeeDao

// employeeDao is the data access object for table cqgf_employee.
// You can define custom methods on it to extend its functionality as you wish.
type employeeDao struct {
	internalEmployeeDao
}

var (
	// Employee is globally public accessible object for table cqgf_employee operations.
	Employee = employeeDao{
		internal.NewEmployeeDao(),
	}
)

// Fill with you ideas below.
