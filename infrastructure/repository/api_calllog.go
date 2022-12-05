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
	ApiCallLogRepository interface {
		Create(ctx context.Context, log *entity.ApiCallLog) error
	}
	apiCallLogRepository struct {
		res    resource.Resource
		client Client
	}
)

func (a *apiCallLogRepository) Create(ctx context.Context, log *entity.ApiCallLog) error {
	if err := a.client.DB(ctx).Create(log).Error; err != nil {
		return err
	}
	return nil
}

func newApiCallLogRepository(res resource.Resource, client Client) ApiCallLogRepository {
	return &apiCallLogRepository{
		res:    res,
		client: client,
	}
}
