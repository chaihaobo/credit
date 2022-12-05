// Package controller
// @author： Boice
// @createTime：2022/11/28 11:09
package controller

import (
	"credit-platform/resource"
	"credit-platform/service"
)

type (
	Controller interface {
		Api() ApiController
	}
	controller struct {
		res resource.Resource
		svc service.Service
		api ApiController
	}
)

func (c controller) Api() ApiController {
	return c.api
}

func New(res resource.Resource, svc service.Service) Controller {

	return &controller{
		res: res,
		svc: svc,
		api: newApiController(res, svc),
	}
}
