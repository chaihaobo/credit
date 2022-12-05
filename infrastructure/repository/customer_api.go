// Package repository
// @author： Boice
// @createTime：2022/11/29 16:54
package repository

import (
	"context"
	"credit-platform/entity"
	"credit-platform/resource"
)

type (
	CustomerApiRepository interface {
		ListCustomerApi(ctx context.Context, customerID int64) ([]*entity.CustomerApi, error)
	}
	customerApiRepository struct {
		res    resource.Resource
		client Client
	}
)

func (c *customerApiRepository) ListCustomerApi(ctx context.Context, customerID int64) ([]*entity.CustomerApi, error) {
	customerApis := make([]*entity.CustomerApi, 0)
	if err := c.client.DB(ctx).Where("customer_id = ?", customerID).Find(&customerApis).Error; err != nil {
		return nil, err
	}
	return customerApis, nil
}

func newCustomerApiRepository(res resource.Resource, c Client) CustomerApiRepository {
	return &customerApiRepository{
		res:    res,
		client: c,
	}
}
