package basic

import (
	"context"
	"fmt"
	"gf_crud/api/basic"
)

type departmentController struct{}

func New() *departmentController {
	return &departmentController{}
}

// 传入什么就响应什么
func (c *departmentController) Test(ctx context.Context, req *basic.DepartmentReq) (res *basic.DepartmentRes, err error) {
	fmt.Println("--------------------", req.Department)
	fmt.Println(req.Department.Id, req.Department.DepartmentName)
	res = &basic.DepartmentRes{
		Department: req.Department,
	}
	return
}
