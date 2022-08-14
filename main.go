package main

import (
	_ "github.com/cool-team-official/cool-admin-go/internal/packed"
	_ "github.com/cool-team-official/cool-admin-go/modules"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/cool-team-official/cool-admin-go/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
