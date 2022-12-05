// Package transport
// @author： Boice
// @createTime：2022/11/29 09:48
package transport

import (
	"context"
	"credit-platform/infrastructure"
	"credit-platform/resource"
	"credit-platform/service"
	"credit-platform/transport/http"
	"go.uber.org/zap"
)

type (
	Transport interface {
		Serve()
		Shutdown()
	}
	transport struct {
		res  resource.Resource
		svc  service.Service
		http http.Http
	}
)

func (t *transport) Shutdown() {
	t.http.Shutdown()
	t.res.Logger().Info(context.Background(), "all server stopped!")
}

func (t *transport) Serve() {
	go t.http.Listen()
	t.res.Logger().Info(context.Background(), "http server started!", zap.String("addr", t.res.Config().Server.Addr))
}

func New(res resource.Resource, svc service.Service, infra infrastructure.Infrastructure) Transport {
	h := http.New(res, svc, infra)
	return &transport{
		res:  res,
		svc:  svc,
		http: h,
	}
}
