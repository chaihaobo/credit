package minivision

import "credit-platform/constant"

var (
	errorCodeMapping = map[int]constant.ServiceError{
		0: constant.ErrInternalServerError,
		// face compare
		2001: constant.ErrApiFaceCompareImage1Required,
		2011: constant.ErrApiFaceCompareImage2Required,
		2003: constant.ErrApiFaceCompareImage1NoFaceDetected,
		2013: constant.ErrApiFaceCompareImage2NoFaceDetected,
		// livingBody
		1700: constant.ErrApiLivingBodyFaceNotExist,
		2800: constant.ErrApiLivingBodyFaceNotExist,
		1004: constant.ErrApiLivingBodyCopyFileFailed,
	}
)

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
	if errData, ok := br.Data.(map[string]any); ok {
		minivisionErrCode := errData["errorCode"].(float64)
		if err, ok := errorCodeMapping[int(minivisionErrCode)]; ok {
			return err
		}
	}
	return constant.ErrInternalServerError
}
