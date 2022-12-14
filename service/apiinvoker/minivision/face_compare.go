package minivision

import (
	"credit-platform/constant"
	"credit-platform/infrastructure"
	"credit-platform/model/api"
	"credit-platform/resource"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type (
	FaceCompareInvoker interface {
		Invoke(c *gin.Context) (any, error)
	}
	faceCompareInvoker struct {
		client Client
		res    resource.Resource
		infra  infrastructure.Infrastructure
	}
)

func (f *faceCompareInvoker) Invoke(c *gin.Context) (any, error) {
	ctx := c.Request.Context()
	req := new(api.FaceCompareRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, err
	}
	minivisionFaceCompareRequest := req.ToMinivisionFaceCompareRequest()
	f.res.Logger().Info(ctx, "minivision face compare request", zap.Any("param", minivisionFaceCompareRequest))
	compareResult, err := f.client.FaceCompare(ctx, minivisionFaceCompareRequest)
	if err != nil {
		f.res.Logger().Error(ctx, "minivision face compare error", zap.Error(err))
		return nil, constant.ErrInternalServerError
	}
	f.res.Logger().Info(ctx, "minivision face compare response", zap.Any("result", compareResult))
	if !compareResult.Normal() {
		return nil, compareResult.Error()
	}
	return compareResult.Data.(float64), nil
}

func NewFaceCompareInvoker(res resource.Resource, infra infrastructure.Infrastructure) FaceCompareInvoker {
	return &faceCompareInvoker{
		client: newClient(res, infra),
		res:    res,
		infra:  infra,
	}
}
