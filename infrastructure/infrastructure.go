// Package infrastructure
// @author： Boice
// @createTime：2022/11/29 10:30
package infrastructure

import (
	"credit-platform/infrastructure/repository"
	"credit-platform/resource"
)

type (
	Infrastructure interface {
		Repository() repository.Repository
	}
	infrastructure struct {
		res  resource.Resource
		repo repository.Repository
	}
)

func (i *infrastructure) Repository() repository.Repository {
	return i.repo
}

func New(res resource.Resource) Infrastructure {
	repo := repository.New(res)
	return &infrastructure{
		res:  res,
		repo: repo,
	}
}
