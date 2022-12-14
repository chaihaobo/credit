package api

import (
	"credit-platform/model/minivision"
)

type LiveBodyRequest struct {
	ImageBase64 []byte `json:"image_base64" binding:"required"`
}

func (r *LiveBodyRequest) ToMinivisionFaceCompareRequest() *minivision.LiveBodyRequest {
	return &minivision.LiveBodyRequest{
		Base64: r.ImageBase64,
	}
}
