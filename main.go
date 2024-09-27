package main

import (
	_ "ipv6Host/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"ipv6Host/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
