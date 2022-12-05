// Package repository
// @author： Boice
// @createTime：2022/11/29 16:14
package repository

import (
	"context"
	"credit-platform/entity"
	"credit-platform/resource"
	"gorm.io/gorm/clause"
)

type (
	CustomerRepository interface {
		GetByLoginName(ctx context.Context, loginName string) (*entity.Customer, error)
		GetByToken(ctx context.Context, token string) (*entity.Customer, error)
		GetByIdForUpdate(ctx context.Context, id int64) (*entity.Customer, error)
		Update(ctx context.Context, customer *entity.Customer) error
	}
	customerRepository struct {
		res    resource.Resource
		client Client
	}
)

func (c *customerRepository) GetByToken(ctx context.Context, token string) (*entity.Customer, error) {
	customer := new(entity.Customer)
	if result := c.client.DB(ctx).Where("token = ?", token).Find(customer); result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	return customer, nil
}

func (c *customerRepository) Update(ctx context.Context, customer *entity.Customer) error {
	if err := c.client.DB(ctx).Updates(customer).Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) GetByIdForUpdate(ctx context.Context, id int64) (*entity.Customer, error) {
	customer := new(entity.Customer)
	if result := c.client.DB(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).Find(customer, id); result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	return customer, nil
}

func (c *customerRepository) GetByLoginName(ctx context.Context, loginName string) (*entity.Customer, error) {
	customer := new(entity.Customer)
	if result := c.client.DB(ctx).Where("login_name = ?", loginName).Find(customer); result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	return customer, nil
}

func newCustomerRepository(res resource.Resource, client Client) CustomerRepository {
	return &customerRepository{
		res:    res,
		client: client,
	}
}
