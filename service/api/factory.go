// Package apiinvoker
// @author： Boice
// @createTime：2022/11/29 18:54
package api

import (
	"credit-platform/constant/enum"
	"credit-platform/infrastructure"
	"credit-platform/resource"
	"credit-platform/service/api/invoker"
	"credit-platform/service/api/invoker/minivision"
)

type (
	InvokerFactory interface {
		Invoker(path enum.ApiPath) invoker.Invoker
	}
	invokerFactory struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
	}
)

func (f *invokerFactory) Invoker(path enum.ApiPath) invoker.Invoker {
	switch path {
	case enum.ApiPathFaceCompare:
		return minivision.NewFaceCompareInvoker(f.res, f.infra)
	case enum.ApiPathLiveBody:
		return minivision.NewLiveBodyInvoker(f.res, f.infra)
	default:
		return nil

	}
}

func NewInvokerFactory(res resource.Resource, infra infrastructure.Infrastructure) InvokerFactory {
	return &invokerFactory{
		res:   res,
		infra: infra,
	}
}
