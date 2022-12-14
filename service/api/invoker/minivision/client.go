package minivision

import (
	"context"
	"credit-platform/infrastructure"
	"credit-platform/model/minivision"
	"credit-platform/resource"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

type (
	Client interface {
		//	FaceCompare 人脸比对
		//	输入两张base64照片,比对相似度
		FaceCompare(ctx context.Context, request *minivision.FaceCompareRequest) (*minivision.BaseResponse, error)

		//	LiveBody 活体检测
		//	输入一张base64照片,检查是否真人拍摄
		LiveBody(ctx context.Context, request *minivision.LiveBodyRequest) (*minivision.BaseResponse, error)
	}
	client struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
		c     *resty.Client
	}
)

func newClient(res resource.Resource, infra infrastructure.Infrastructure) Client {
	return &client{
		res:   res,
		infra: infra,
		c:     resty.New(),
	}
}

func (c *client) FaceCompare(ctx context.Context, request *minivision.FaceCompareRequest) (*minivision.BaseResponse, error) {
	baseRequest, token := c.baseParam()
	request.BaseRequest = baseRequest
	result := new(minivision.BaseResponse)
	_, err := c.c.R().EnableTrace().
		SetContext(ctx).
		SetBody(request).
		SetHeader("Token", token).
		SetResult(result).
		Post(c.res.Config().Api.Minivision.Url + "/api/v3/compare")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) LiveBody(ctx context.Context, request *minivision.LiveBodyRequest) (*minivision.BaseResponse, error) {
	baseRequest, token := c.baseParam()
	request.BaseRequest = baseRequest
	result := new(minivision.BaseResponse)
	_, err := c.c.R().
		SetContext(ctx).
		SetBody(request).
		SetHeader("Token", token).
		SetResult(result).
		Post(c.res.Config().Api.Minivision.Url + "/api/v3/livingBody")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) baseParam() (baseRequest minivision.BaseRequest, token string) {
	appKey := c.res.Config().Api.Minivision.AppKey
	appSecret := c.res.Config().Api.Minivision.AppSecret
	timestamp := strconv.Itoa(int(time.Now().UnixMilli()))
	unEncToken := fmt.Sprintf("%s{%s:%s}", appKey, appSecret, timestamp)
	hash := md5.New()
	hash.Write([]byte(unEncToken))
	token = hex.EncodeToString(hash.Sum(nil))
	baseRequest = minivision.BaseRequest{
		AppKey:    appKey,
		Timestamp: timestamp,
	}
	return
}
