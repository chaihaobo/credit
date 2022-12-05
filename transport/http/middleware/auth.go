// Package middlewire
// @author： Boice
// @createTime：2022/11/29 15:11
package middleware

import (
	"context"
	"credit-platform/constant"
	"credit-platform/entity"
	"credit-platform/infrastructure"
	"credit-platform/resource"
	"credit-platform/util"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const (
	headerToken = "X-Credit-Token"
)

//	AuthMiddleware 请求验签中间件
func AuthMiddleware(res resource.Resource, infra infrastructure.Infrastructure) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.GetHeader(headerToken)

		if username == "" {
			err := constant.ErrAuthTokenRequired
			res.Logger().Error(ctx, "鉴权token为空", zap.Error(err))
			util.HttpError(ctx, err)
			return
		}
		customer, err := loadCustomer(ctx.Request.Context(), res, infra, username)
		if err != nil {
			util.HttpError(ctx, err)
			return
		}
		if customer == nil {
			util.HttpError(ctx, constant.ErrAuthTokenInvalid)
			return
		}
		// 保存客户到上下文
		request := ctx.Request
		ctx.Request = request.WithContext(context.WithValue(request.Context(), constant.CtxCustomer, customer))
		ctx.Next()

	}
}

//	loadCustomer 根据token加载客户以及api信息
func loadCustomer(ctx context.Context, res resource.Resource, infra infrastructure.Infrastructure, token string) (*entity.Customer, error) {
	customer, err := infra.Repository().Customer().GetByToken(ctx, token)
	if err != nil {
		res.Logger().Error(ctx, "auth middleware get customer error", zap.Error(err))
		return nil, err
	}
	if customer == nil {
		res.Logger().Error(ctx, "auth middleware get customer nil", zap.Error(constant.ErrAuthTokenInvalid))
		return nil, constant.ErrAuthTokenInvalid
	}
	// 获取用户api权限信息
	customerApis, err := infra.Repository().CustomerApi().ListCustomerApi(ctx, customer.ID)
	if err != nil {
		res.Logger().Error(ctx, "auth middleware get customer apis", zap.Error(err))
		return nil, err
	}
	apiList := lo.Map[*entity.CustomerApi, *entity.Api](customerApis, func(customerApi *entity.CustomerApi, _ int) *entity.Api {
		api, _ := infra.Repository().Api().GetByID(ctx, customerApi.ID)
		return api
	})
	customer.ApiList = apiList
	return customer, nil

}
