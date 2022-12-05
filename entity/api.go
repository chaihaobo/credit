// Package entity
// @author： Boice
// @createTime：2022/11/29 15:06
package entity

import "github.com/shopspring/decimal"

type Api struct {
	BaseEntity
	Name  string          `gorm:"name"`
	Path  string          `gorm:"path"`
	Price decimal.Decimal `gorm:"price"`
}

func (*Api) TableName() string {
	return "t_api"
}
