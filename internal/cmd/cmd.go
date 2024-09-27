package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"ipv6Host/internal/controller/hello"
	"ipv6Host/internal/controller/host"
	"ipv6Host/internal/controller/admin"
	"ipv6Host/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 启动管理后台gtoken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				g.Log().Error(ctx, err)
			}
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					hello.New(),
					admin.New().Create,
					host.New(),
					// login.New(),
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfAdminToken.Middleware(ctx, group) // 使用 gctx.New() 创建新的上下文
					if err != nil {
						panic(err)
					}
					group.Bind(
						
						admin.New().List,
						admin.New().Update,
						admin.New().Delete,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
