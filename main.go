package main

import (
	_ "gf_crud/internal/logic"
	_ "gf_crud/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"gf_crud/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
