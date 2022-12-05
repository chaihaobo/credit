// Package util
// @author： Boice
// @createTime：2022/11/29 16:20
package util

import (
	"credit-platform/constant"
	"credit-platform/model"
	"github.com/gin-gonic/gin"
)

func HttpError(ctx *gin.Context, err error) {
	var serviceError constant.ServiceError
	switch err.(type) {
	case constant.ServiceError:
		serviceError = err.(constant.ServiceError)
	default:
		serviceError = constant.ErrInternalServerError
	}
	ctx.JSON(serviceError.HttpCode, model.NewBaseResponseFromError(ctx.Request.Context(), err))
	ctx.Abort()
}
