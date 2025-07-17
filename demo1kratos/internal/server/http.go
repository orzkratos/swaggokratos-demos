package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/demokratos/demo1kratos"
	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
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
			ContentData: demo1kratos.GetOpenapiContent("demo1kratos-title"),
		},
	})

	zapLog.SUG.Infoln("[DOC]", "(http://127.0.0.1:"+swaggokratos.MustGetPortNum(c.Http.Addr)+"/doc/swagger/a/index.html)")
	zapLog.SUG.Infoln("接口文档添加成功")
}
