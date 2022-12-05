// Package entity
// @author： Boice
// @createTime：2022/11/29 15:08
package entity

type CustomerApi struct {
	BaseEntity
	CustomerId int64 `gorm:"customer_id"`
	ApiId      int64 `gorm:"api_id"`
}

func (*CustomerApi) TableName() string {
	return "t_customer_api"
}
