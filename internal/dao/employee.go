// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_crud/internal/dao/internal"
)

// internalEmployeeDao is internal type for wrapping internal DAO implements.
type internalEmployeeDao = *internal.EmployeeDao

// employeeDao is the data access object for table employee.
// You can define custom methods on it to extend its functionality as you wish.
type employeeDao struct {
	internalEmployeeDao
}

var (
	// Employee is globally public accessible object for table employee operations.
	Employee = employeeDao{
		internal.NewEmployeeDao(),
	}
)

// Fill with you ideas below.
