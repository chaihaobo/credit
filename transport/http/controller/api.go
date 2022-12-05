// Package controller
// @author： Boice
// @createTime：2022/11/29 12:19
package controller

import (
	"credit-platform/constant/enum"
	"credit-platform/resource"
	"credit-platform/service"
	"github.com/gin-gonic/gin"
)

type (
	ApiController interface {
		Invoke(ctx *gin.Context) (any, error)
	}
	apiController struct {
		res resource.Resource
		svc service.Service
	}
)

func (a apiController) Invoke(ctx *gin.Context) (any, error) {
	apiPath := ctx.Param("apiPath")
	result, err := a.svc.Api().Invoke(ctx, enum.ParseApiPath(apiPath))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func newApiController(res resource.Resource, svc service.Service) ApiController {
	return &apiController{
		res: res,
		svc: svc,
	}
}
