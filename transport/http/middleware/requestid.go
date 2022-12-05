// Package middlewire
// @author： Boice
// @createTime：2022/11/29 14:27
package middleware

import (
	"context"
	"credit-platform/constant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request
		ctx.Request = request.WithContext(context.WithValue(request.Context(), constant.CtxRequestID, uuid.New().String()))
		ctx.Next()
	}
}
