package basic

import (
	"context"
	"database/sql"
	"fmt"
	"gf_crud/internal/dao"
	"gf_crud/internal/model"
	"gf_crud/internal/model/do"
	"gf_crud/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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

// 新增部门
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

// 删除部门
func (s *sDepartment) DeleteDepartment(ctx context.Context, in model.DepartmentInput) (flag bool, err error) {
	idList := in.IdList
	fmt.Println("idList: ", idList)
	if len(idList) <= 0 {
		return false, gerror.NewCode(gcode.CodeInvalidParameter, "id不能空")
	}
	r, err2 := dao.Department.Ctx(ctx).WhereIn("id", idList).Delete()
	fmt.Println(r, err2)
	count, _ := r.RowsAffected()
	return count > 0, nil
}
