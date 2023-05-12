package basic

import (
	"context"
	"fmt"
	v1 "gf_crud/api/basic/v1"
)

type departmentController struct{}

func New() *departmentController {
	return &departmentController{}
}

// 传入什么就响应什么
func (c *departmentController) Test(ctx context.Context, req *v1.DepartmentReq) (res *v1.DepartmentRes, err error) {
	fmt.Println("--------------------", req.Department)
	fmt.Println(req.Department.Id, req.Department.DepartmentName)
	res = &v1.DepartmentRes{
		Department: req.Department,
	}
	return
}
