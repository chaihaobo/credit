// Package constant
// @author： Boice
// @createTime：2022/11/29 12:37
package constant

import (
	"fmt"
	"net/http"
)

var (
	Success                = ServiceError{HttpCode: http.StatusOK, Code: "0000", Message: "success"}
	ErrInternalServerError = ServiceError{HttpCode: http.StatusInternalServerError, Code: "9999", Message: "internal error"}

	ErrAuthTokenRequired                  = ServiceError{HttpCode: http.StatusUnauthorized, Code: "1001", Message: "token required"}
	ErrAuthTokenInvalid                   = ServiceError{HttpCode: http.StatusUnauthorized, Code: "1002", Message: "auth token invalid"}
	ErrUserPasswordInvalid                = ServiceError{HttpCode: http.StatusUnauthorized, Code: "1003", Message: "user password invalid"}
	ErrApiInvalid                         = ServiceError{HttpCode: http.StatusBadRequest, Code: "1004", Message: "user password invalid"}
	ErrCustomerNotOpenApi                 = ServiceError{HttpCode: http.StatusBadRequest, Code: "1005", Message: "customer not open this api"}
	ErrCustomerBalanceInsufficient        = ServiceError{HttpCode: http.StatusBadRequest, Code: "1006", Message: "customer balance insufficient"}
	ErrApiNotImplement                    = ServiceError{HttpCode: http.StatusBadRequest, Code: "1007", Message: "api not implement"}
	ErrApiFaceCompareImage1Required       = ServiceError{HttpCode: http.StatusBadRequest, Code: "1008", Message: "api face compare image1 required"}
	ErrApiFaceCompareImage2Required       = ServiceError{HttpCode: http.StatusBadRequest, Code: "1009", Message: "api face compare image2 required"}
	ErrApiFaceCompareImage1NoFaceDetected = ServiceError{HttpCode: http.StatusBadRequest, Code: "1010", Message: "api face compare image1 no face detected"}
	ErrApiFaceCompareImage2NoFaceDetected = ServiceError{HttpCode: http.StatusBadRequest, Code: "1011", Message: "api face compare image2 no face detected"}
	ErrApiLivingBodyFaceNotExist          = ServiceError{HttpCode: http.StatusBadRequest, Code: "1012", Message: "api living body face not exist"}
	ErrApiLivingBodyCopyFileFailed        = ServiceError{HttpCode: http.StatusInternalServerError, Code: "1013", Message: "api living body copy file failed"}
)

type (
	ServiceError struct {
		HttpCode int
		Code     string
		Message  string
	}
)

func (s ServiceError) Error() string {
	return fmt.Sprintf("http code: %d, code: %s, message: %s", s.HttpCode, s.Code, s.Message)
}
