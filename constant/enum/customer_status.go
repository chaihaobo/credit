// Package enum
// @author： Boice
// @createTime：2022/11/29 15:05
package enum

const (
	CustomerStatusActive   CustomerStatus = "active"
	CustomerStatusInActive                = "in_active"
)

type CustomerStatus string
