package basic

import (
	"context"
	"fmt"
	"gf_crud/api/basic"
	"gf_crud/internal/model"
	"gf_crud/internal/service"
)

type departmentController struct{}

func New() *departmentController {
	return &departmentController{}
}

// 传入什么就响应什么
func (c *departmentController) Test(ctx context.Context, req *basic.DepartmentReq) (res *basic.DepartmentRes, err error) {
	fmt.Println("--------------------", req.Department)
	fmt.Println(req.Department.Id, req.Department.DepartmentName)
	dataInput := model.DepartmentInput{
		Id:             req.Department.Id,
		DepartmentName: req.DepartmentName,
	}
	dataOutput, err := service.Department().Test(ctx, dataInput)
	res = &basic.DepartmentRes{
		Department: basic.Department{
			Id:             dataOutput.Id,
			DepartmentName: dataOutput.DepartmentName,
		},
	}
	fmt.Println("controller---", &res)
	return
}
