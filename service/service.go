// Package service
// @author： Boice
// @createTime：2022/11/28 11:18
package service

import (
	"credit-platform/infrastructure"
	"credit-platform/resource"
)

type (
	Service interface {
		Api() ApiService
	}
	service struct {
		infra infrastructure.Infrastructure
		res   resource.Resource
		api   ApiService
	}
)

func (s *service) Api() ApiService {
	return s.api
}

func New(infra infrastructure.Infrastructure, res resource.Resource) Service {
	return &service{
		infra: infra,
		res:   res,
		api:   newApiService(infra, res),
	}
}
