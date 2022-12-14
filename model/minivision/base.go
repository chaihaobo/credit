package minivision

type (
	BaseRequest struct {
		AppKey    string `json:"appKey"`
		Timestamp string `json:"timestamp"`
	}
	BaseResponse struct {
		RequestId string `json:"requestId"`
		TimeUsed  int64  `json:"timeUsed"`
		Status    int    `json:"status"`
		Data      any    `json:"data"`
	}

	Error struct {
		ErrorCode    int    `json:"errorCode"`
		Timestamp    int64  `json:"timestamp"`
		ErrorMessage string `json:"errorMessage"`
		Path         string `json:"path"`
	}
)

func (br *BaseResponse) Normal() bool {
	return br.Status == 0
}

func (br *BaseResponse) Error() error {
	if br.Normal() {
		return nil
	}

	return nil
}
