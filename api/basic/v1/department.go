package v1

import (
	"gf_crud/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DepartmentReq struct {
	g.Meta `path:"/department/test" method:"post" tags:"department查询" summary:"department查询"`
	*entity.Department
}

type DepartmentRes struct {
	*entity.Department
}
