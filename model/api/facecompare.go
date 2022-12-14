package api

import (
	"credit-platform/model/minivision"
)

type FaceCompareRequest struct {
	Image1Base64 []byte `json:"image1_base64" binding:"required"`
	Image2Base64 []byte `json:"image2_base64" binding:"required"`
}

func (r *FaceCompareRequest) ToMinivisionFaceCompareRequest() *minivision.FaceCompareRequest {
	return &minivision.FaceCompareRequest{
		Image1Base64: r.Image1Base64,
		Image2Base64: r.Image2Base64,
	}
}
