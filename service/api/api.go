// Package service
// @author： Boice
// @createTime：2022/11/29 17:56
package api

import (
	"context"
	"credit-platform/constant"
	"credit-platform/constant/enum"
	"credit-platform/entity"
	"credit-platform/infrastructure"
	"credit-platform/resource"
	"credit-platform/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type (
	Service interface {
		Invoke(c *gin.Context, apiPath enum.ApiPath) (any, error)
	}
	service struct {
		res               resource.Resource
		infra             infrastructure.Infrastructure
		apiInvokerFactory InvokerFactory
	}
)

func (a *service) Invoke(c *gin.Context, apiPath enum.ApiPath) (any, error) {
	ctx := c.Request.Context()
	if apiPath == enum.ApiPathUnKnow {
		return nil, constant.ErrApiInvalid
	}
	customer := util.ContextCustomer(c.Request.Context())
	// 判断用户是否开通此api
	if !customer.HasApi(apiPath) {
		return nil, constant.ErrCustomerNotOpenApi
	}
	// 判断这个api是否实现
	apiInvoker := a.apiInvokerFactory.Invoker(apiPath)
	if apiInvoker == nil {
		return nil, constant.ErrApiNotImplement
	}
	api := customer.GetApi(apiPath)
	ctx = a.infra.Repository().Begin(ctx)
	defer a.infra.Repository().Rollback(ctx)
	// 判断用户余额是否充足
	realCustomer, err := a.infra.Repository().Customer().GetByIdForUpdate(ctx, customer.ID)
	if err != nil {
		a.res.Logger().Error(ctx, "api service get customer by id for update error", zap.Error(err))
		return nil, constant.ErrInternalServerError
	}
	if realCustomer.Balance.LessThan(api.Price) {
		a.res.Logger().Error(ctx, "customer balance insufficient")
		return nil, constant.ErrCustomerBalanceInsufficient
	}
	// 执行 api
	invoke, err := apiInvoker.Invoke(c)
	// 保存api调用日志
	a.saveApiInvokeLog(ctx, customer, api, err)
	if err != nil {
		a.infra.Repository().Commit(ctx)
		return nil, err
	}
	//	扣减用户余额
	realCustomer.DeductionBalance(api.Price)
	err = a.infra.Repository().Customer().Update(ctx, realCustomer)
	if err != nil {
		a.res.Logger().Error(ctx, "api service update customer balance error", zap.Error(err))
		return nil, constant.ErrInternalServerError
	}
	a.infra.Repository().Commit(ctx)
	return invoke, nil
}

//	saveApiInvokeLog 保存api调用日志
func (a *service) saveApiInvokeLog(ctx context.Context, customer *entity.Customer, api *entity.Api, err error) {
	log := entity.NewApiCallLog(util.ContextRequestID(ctx), customer, api, err)
	err = a.infra.Repository().ApiCallLog().Create(ctx, log)
	if err != nil {
		a.res.Logger().Error(ctx, "api service save api invoke log error", zap.Error(err))
	}
}

func NewService(infra infrastructure.Infrastructure, res resource.Resource) Service {
	return &service{
		res:               res,
		infra:             infra,
		apiInvokerFactory: NewInvokerFactory(res, infra),
	}
}
