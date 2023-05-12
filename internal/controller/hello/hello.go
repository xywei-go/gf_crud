package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_crud/api/hello/v1"
)

type HelloController struct{}

func New() *HelloController {
	return &HelloController{}
}

func (c *HelloController) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
