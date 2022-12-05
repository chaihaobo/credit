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

	ErrAuthTokenRequired           = ServiceError{HttpCode: http.StatusUnauthorized, Code: "1001", Message: "token required"}
	ErrAuthTokenInvalid            = ServiceError{HttpCode: http.StatusUnauthorized, Code: "1002", Message: "auth token invalid"}
	ErrUserPasswordInvalid         = ServiceError{HttpCode: http.StatusUnauthorized, Code: "1003", Message: "user password invalid"}
	ErrApiInvalid                  = ServiceError{HttpCode: http.StatusBadRequest, Code: "1004", Message: "user password invalid"}
	ErrCustomerNotOpenApi          = ServiceError{HttpCode: http.StatusBadRequest, Code: "1005", Message: "customer not open this api"}
	ErrCustomerBalanceInsufficient = ServiceError{HttpCode: http.StatusBadRequest, Code: "1006", Message: "customer balance insufficient"}
	ErrApiNotImplement             = ServiceError{HttpCode: http.StatusBadRequest, Code: "1007", Message: "api not implement"}
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
