package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/demokratos/demo2kratos"
	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
	"github.com/orzkratos/demokratos/demo2kratos/internal/conf"
	"github.com/orzkratos/demokratos/demo2kratos/internal/service"
	"github.com/orzkratos/swaggokratos"
	"github.com/orzkratos/swaggokratos/swaggogin"
	"github.com/orzkratos/zapkratos"
	"github.com/yyle88/zaplog"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	serveSwaggerHttpDocument(c, srv)
	return srv
}

func serveSwaggerHttpDocument(c *conf.Server, srv *http.Server) {
	zapKratos := zapkratos.NewZapKratos(zaplog.LOGGER, zapkratos.NewOptions())
	zapLog := zapKratos.SubZap()
	zapLog.SUG.Infoln("准备添加接口文档")

	swaggokratos.RegisterSwaggoHTTPServer(srv, "/doc/", []*swaggogin.Param{
		{
			SwaggerPath: "/swagger/a/*any",
			ExplorePath: "/abc/openapi-a.yaml",
			ContentData: demo2kratos.GetOpenapiContent("自定义的文档标题"),
		},
	})

	zapLog.SUG.Infoln("[DOC]", "(http://127.0.0.1:"+swaggokratos.MustGetPortNum(c.Http.Addr)+"/doc/swagger/a/index.html)")
	zapLog.SUG.Infoln("接口文档添加成功")
}
