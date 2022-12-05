// Package entity
// @author： Boice
// @createTime：2022/11/29 14:55
package entity

import (
	"credit-platform/constant/enum"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type Customer struct {
	BaseEntity
	Name      string              `gorm:"name"`
	Phone     string              `gorm:"phone"`
	Email     string              `gorm:"email"`
	LoginName string              `gorm:"login_name"`
	Password  string              `gorm:"password"`
	Balance   decimal.Decimal     `gorm:"balance"`
	Status    enum.CustomerStatus `gorm:"status"`
	ApiList   []*Api              `gorm:"-"`
}

func (c *Customer) HasApi(apiPath enum.ApiPath) bool {
	return lo.ContainsBy[*Api](c.ApiList, func(api *Api) bool {
		return api.Path == string(apiPath)
	})
}

func (*Customer) TableName() string {
	return "t_customer"
}

//	GetApi 根据api path获取用户api
func (c *Customer) GetApi(path enum.ApiPath) *Api {
	return lo.FindOrElse[*Api](c.ApiList, nil, func(api *Api) bool {
		return api.Path == string(path)
	})
}

func (c *Customer) DeductionBalance(price decimal.Decimal) {
	c.Balance = c.Balance.Sub(price)
}
