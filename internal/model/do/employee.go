// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package do

import(
"github.com/gogf/gf/v2/frame/g"
)

// Employee is the golang structure of table employee for DAO operations like Where/Data.
type Employee struct {
g.Meta `orm:"table:employee, do:true"`
    Id interface{} // 员工ID      
    EmployeeName interface{} // 员工名称    
    Age interface{} // 员工年纪    
    DepartmentId interface{} // 员工部门ID  
}