// Package repository
// @author： Boice
// @createTime：2022/11/29 16:14
package repository

import (
	"context"
	"credit-platform/entity"
	"credit-platform/resource"
)

type (
	ApiRepository interface {
		GetByID(ctx context.Context, id int64) (*entity.Api, error)
	}
	apiRepository struct {
		res    resource.Resource
		client Client
	}
)

func (a *apiRepository) GetByID(ctx context.Context, id int64) (*entity.Api, error) {
	api := new(entity.Api)
	if err := a.client.DB(ctx).Find(api, id).Error; err != nil {
		return nil, err
	}
	return api, nil
}

func newApiRepository(res resource.Resource, client Client) ApiRepository {
	return &apiRepository{
		res:    res,
		client: client,
	}
}
