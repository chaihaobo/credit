package minivision

import (
	"credit-platform/constant"
	"credit-platform/infrastructure"
	"credit-platform/model/api"
	"credit-platform/resource"
	"credit-platform/service/api/invoker"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type (
	liveBodyInvoker struct {
		client Client
		res    resource.Resource
		infra  infrastructure.Infrastructure
	}
)

func (f *liveBodyInvoker) Invoke(c *gin.Context) (any, error) {
	ctx := c.Request.Context()
	req := new(api.LiveBodyRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, err
	}
	minivisionLiveBodyRequest := req.ToMinivisionFaceCompareRequest()
	f.res.Logger().Info(ctx, "minivision face compare request")
	compareResult, err := f.client.LiveBody(ctx, minivisionLiveBodyRequest)
	if err != nil {
		f.res.Logger().Error(ctx, "minivision liveBody error", zap.Error(err))
		return nil, constant.ErrInternalServerError
	}
	f.res.Logger().Info(ctx, "minivision liveBody response", zap.Any("result", compareResult))
	if !compareResult.Normal() {
		return nil, compareResult.Error()
	}
	return compareResult.Data.(bool), nil
}

func NewLiveBodyInvoker(res resource.Resource, infra infrastructure.Infrastructure) invoker.Invoker {
	return &liveBodyInvoker{
		client: newClient(res, infra),
		res:    res,
		infra:  infra,
	}
}
