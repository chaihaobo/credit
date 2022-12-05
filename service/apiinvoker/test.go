// Package apiinvoker
// @author： Boice
// @createTime：2022/11/30 09:43
package apiinvoker

import (
	"credit-platform/infrastructure"
	"credit-platform/resource"
	"github.com/gin-gonic/gin"
)

type (
	testInvoker struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
	}
)

func (t *testInvoker) Invoke(c *gin.Context) (any, error) {
	return "this is a test api", nil
}

func newTestInvoker(res resource.Resource, infra infrastructure.Infrastructure) Invoker {
	return &testInvoker{
		res:   res,
		infra: infra,
	}
}
