// Package apiinvoker
// @author： Boice
// @createTime：2022/11/29 18:54
package apiinvoker

import (
	"credit-platform/constant/enum"
	"credit-platform/infrastructure"
	"credit-platform/resource"
)

type (
	Factory interface {
		Invoker(path enum.ApiPath) Invoker
	}
	factory struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
	}
)

func (f *factory) Invoker(path enum.ApiPath) Invoker {
	switch path {
	case enum.ApiPathTest:
		return newTestInvoker(f.res, f.infra)
	default:
		return nil

	}
}

func NewFactory(res resource.Resource, infra infrastructure.Infrastructure) Factory {
	return &factory{
		res:   res,
		infra: infra,
	}
}
