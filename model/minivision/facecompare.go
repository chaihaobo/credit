package minivision

type (
	FaceCompareRequest struct {
		BaseRequest
		Image1Base64 []byte `json:"image1Base64"`
		Image2Base64 []byte `json:"image2Base64"`
	}
)
