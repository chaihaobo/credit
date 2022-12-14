// Package service
// @author： Boice
// @createTime：2022/11/28 11:18
package service

import (
	"credit-platform/infrastructure"
	"credit-platform/resource"
	"credit-platform/service/api"
)

type (
	Service interface {
		Api() api.Service
	}
	service struct {
		infra infrastructure.Infrastructure
		res   resource.Resource
		api   api.Service
	}
)

func (s *service) Api() api.Service {
	return s.api
}

func New(infra infrastructure.Infrastructure, res resource.Resource) Service {
	return &service{
		infra: infra,
		res:   res,
		api:   api.NewService(infra, res),
	}
}
