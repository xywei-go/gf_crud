package basic

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DepartmentReq struct {
	g.Meta `path:"/department/test" method:"post" tags:"department test" summary:"department test2"`
	Department
}

type DepartmentRes struct {
	Department
}

type Department struct {
	Id             int    `json:"id"             description:"部门ID"`
	DepartmentName string `json:"departmentName" description:"部门名称"`
}

type DepartmentCreateReq struct {
	g.Meta         `path:"/department/create" method:"post" tags:"新增部门1" summary:"新增部门11"`
	DepartmentName string `json:"departmentName" dc:"departmentName"`
}

type DepartmentCreateRes struct {
	Result bool `json:"result" dc:"result"`
}

type DepartmentDeleteReq struct {
	g.Meta `path:"/department/delete" method:"post" tags:"根据ID删除部门" summary:"根据ID删除部门，支持批量"`
	IdList []int `json:"IdList"             description:"部门ID"`
}

type DepartmentDeleteRes struct {
	Result bool `json:"result" dc:"result"`
}
