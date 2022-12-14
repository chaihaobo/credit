// Package http
// @author： Boice
// @createTime：2022/11/29 09:48
package http

import (
	"context"
	"credit-platform/constant"
	"credit-platform/infrastructure"
	"credit-platform/model"
	"credit-platform/resource"
	"credit-platform/service"
	"credit-platform/transport/http/controller"
	"credit-platform/transport/http/middleware"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	netHttp "net/http"
	"time"
)

type (
	Http interface {
		Listen()
		Shutdown()
	}
	http struct {
		res        resource.Resource
		ctrl       controller.Controller
		svc        service.Service
		infra      infrastructure.Infrastructure
		g          *gin.Engine
		server     *netHttp.Server
		translator ut.Translator
	}
)

func (h *http) Listen() {
	// apply router
	h.g.Use(middleware.AuthMiddleware(h.res, h.infra))
	h.g.Use(middleware.RequestIDMiddleWare())
	h.g.POST("/api/:apiPath", h.wrapperController(h.ctrl.Api().Invoke))

	err := h.server.ListenAndServe()
	if err != nil {
		h.res.Logger().Panic(context.Background(), "http listen error", zap.Error(err))
	}
}

func (h *http) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := h.server.Shutdown(ctx)
	if err != nil {
		h.res.Logger().Panic(context.Background(), "http shutdown error", zap.Error(err))
	}
}

func (h *http) wrapperController(fun func(*gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var serviceError = constant.Success
		got, err := fun(c)
		if err != nil {
			h.res.Logger().Error(ctx, "http request error", zap.Error(err))
			switch err.(type) {
			case constant.ServiceError:
				serviceError = err.(constant.ServiceError)
			case validator.ValidationErrors:
				validationErrors := err.(validator.ValidationErrors)
				serviceError = constant.ServiceError{
					HttpCode: netHttp.StatusBadRequest,
					Code:     "4001",
					Message:  validationErrors[0].Translate(h.translator),
				}

			default:
				serviceError = constant.ErrInternalServerError
			}
		}
		baseResponse := model.NewBaseResponse[any](ctx, serviceError.Code, serviceError.Message, got)
		baseResponse.Data = got
		c.JSON(serviceError.HttpCode, baseResponse)
	}
}

func New(res resource.Resource, svc service.Service, infra infrastructure.Infrastructure) Http {
	c := controller.New(res, svc)
	gin.SetMode(res.Config().Server.Model)
	g := gin.New()
	g.Use(ginzap.GinzapWithConfig(res.Logger().Zap(), &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		Context: func(c *gin.Context) []zapcore.Field {
			field := make([]zapcore.Field, 0)
			requestID, ok := c.Request.Context().Value(constant.CtxRequestID).(string)
			if ok {
				field = append(field, zap.String("request_id", requestID))
			}
			return field
		},
	}),
		ginzap.RecoveryWithZap(res.Logger().Zap(), true))
	server := &netHttp.Server{
		Addr:    res.Config().Server.Addr,
		Handler: g,
	}

	trans, _ := ut.New(en.New()).GetTranslator("en")
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enTranslations.RegisterDefaultTranslations(validate, trans)
	}

	return &http{
		res:        res,
		ctrl:       c,
		svc:        svc,
		infra:      infra,
		g:          g,
		server:     server,
		translator: trans,
	}
}
