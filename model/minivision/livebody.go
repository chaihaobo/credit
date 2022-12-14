package minivision

type (
	LiveBodyRequest struct {
		BaseRequest
		Base64 []byte `json:"base64"`
	}
)
