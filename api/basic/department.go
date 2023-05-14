package basic

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DepartmentReq struct {
	g.Meta `path:"/department/test" method:"post" tags:"department查询" summary:"department查询"`
	Department
}

type DepartmentRes struct {
	Department
}

type Department struct {
	Id             int    `json:"id"             description:"部门ID"`
	DepartmentName string `json:"departmentName" description:"部门名称"`
}
