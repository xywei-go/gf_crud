package basic

import (
	"context"
	"database/sql"
	"fmt"
	"gf_crud/internal/dao"
	"gf_crud/internal/model"
	"gf_crud/internal/model/do"
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
		Id:             input.Id,
		DepartmentName: input.DepartmentName,
	}
	fmt.Println("logic-----", &output)
	return output, err
}

func (s *sDepartment) CreateDepartment(ctx context.Context, in model.DepartmentInput) (flag bool, err error) {
	var result sql.Result
	result, _ = dao.Department.Ctx(ctx).Data(do.Department{
		Id:             in.Id,
		DepartmentName: in.DepartmentName,
	}).OnDuplicate("department_name").Save()
	count, _ := result.RowsAffected()
	flag = (count != 0)
	return flag, nil
}
