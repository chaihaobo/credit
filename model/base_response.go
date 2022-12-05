// Package model
// @author： Boice
// @createTime：2022/11/29 14:05
package model

import (
	"context"
	"credit-platform/constant"
)

type BaseResponse[T any] struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
	Data      T      `json:"data"`
}

func NewBaseResponse[T any](ctx context.Context, code string, message string, data T) *BaseResponse[T] {
	return &BaseResponse[T]{
		Code:    code,
		Message: message,
		RequestID: func() string {
			value := ctx.Value(constant.CtxRequestID)
			if value == nil {
				return ""
			}
			return value.(string)
		}(),
		Data: data,
	}
}

func NewBaseResponseFromError(ctx context.Context, err error) *BaseResponse[any] {
	var serviceError constant.ServiceError
	switch err.(type) {
	case constant.ServiceError:
		serviceError = err.(constant.ServiceError)
	default:
		serviceError = constant.ErrInternalServerError
	}
	return NewBaseResponse[any](ctx, serviceError.Code, serviceError.Message, nil)
}
