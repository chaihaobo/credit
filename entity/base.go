// Package entity
// @author： Boice
// @createTime：2022/11/29 15:02
package entity

import "time"

type BaseEntity struct {
	ID         int64     `gorm:"id"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
}
