// Package util
// @author： Boice
// @createTime：2022/11/29 18:09
package util

import (
	"context"
	"credit-platform/constant"
	"credit-platform/entity"
)

func ContextCustomer(ctx context.Context) *entity.Customer {
	value := ctx.Value(constant.CtxCustomer)
	if value != nil {
		return value.(*entity.Customer)
	}
	return nil
}

func ContextRequestID(ctx context.Context) string {
	value := ctx.Value(constant.CtxRequestID)
	if value != nil {
		return value.(string)
	}
	return ""
}
