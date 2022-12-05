// Package entity
// @author： Boice
// @createTime：2022/11/29 15:08
package entity

import (
	"credit-platform/constant/enum"
	"github.com/shopspring/decimal"
	"time"
)

type ApiCallLog struct {
	BaseEntity
	RequestID    string             `gorm:"request_id"`
	CustomerID   int64              `gorm:"customer_id"`
	CustomerName string             `gorm:"customer_name"`
	CallPrice    decimal.Decimal    `gorm:"call_price"`
	ApiID        int64              `gorm:"api_id"`
	ApiName      string             `gorm:"api_name"`
	Status       enum.ApiCallStatus `gorm:"status"`
	FailReason   *string            `gorm:"fail_reason"`
}

func (*ApiCallLog) TableName() string {
	return "t_api_call_log"
}

func NewApiCallLog(requestID string, customer *Customer, api *Api, err error) *ApiCallLog {
	return &ApiCallLog{
		BaseEntity: BaseEntity{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
		RequestID:    requestID,
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		CallPrice: func() decimal.Decimal {
			if err == nil {
				return api.Price
			}
			return decimal.Zero
		}(),
		ApiID:   api.ID,
		ApiName: api.Name,
		Status: func() enum.ApiCallStatus {
			if err != nil {
				return enum.ApiCallStatusFail
			}
			return enum.ApiCallStatusSuccess
		}(),
		FailReason: func() *string {
			if err != nil {
				failReason := err.Error()
				return &failReason
			}
			return nil
		}(),
	}
}
