package basic

import (
	"context"
	"fmt"
	"gf_crud/internal/model"
	"gf_crud/internal/service"
)

type sDepartment struct{}

func init() {
	service.RegisterDepartment(New())
}

func New() service.IDepartment {
	return &sDepartment{}
}

func (s *sDepartment) Test(ctx context.Context, input model.DepartmentInput) (output *model.DepartmentOutput, err error) {
	fmt.Println("---hello, department service test---")
	output = &model.DepartmentOutput{
		Id: input.Id,
		DepartmentName: input.DepartmentName,
	}
	fmt.Println("logic-----", &output)
	return output, err
}
