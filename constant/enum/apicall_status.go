// Package enum
// @author： Boice
// @createTime：2022/11/29 15:10
package enum

const (
	ApiCallStatusSuccess ApiCallStatus = "success"
	ApiCallStatusFail    ApiCallStatus = "fail"
)

type ApiCallStatus string
