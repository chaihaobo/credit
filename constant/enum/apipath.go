// Package enum
// @author： Boice
// @createTime：2022/11/29 17:42
package enum

const (
	ApiPathUnKnow      ApiPath = "unknow"
	ApiPathFaceCompare ApiPath = "faceCompare"
	ApiPathLiveBody    ApiPath = "liveBody"
)

var apiPathMap = map[string]ApiPath{
	string(ApiPathUnKnow):      ApiPathUnKnow,
	string(ApiPathFaceCompare): ApiPathFaceCompare,
	string(ApiPathLiveBody):    ApiPathLiveBody,
}

type ApiPath string

func ParseApiPath(path string) ApiPath {
	if apiPath, ok := apiPathMap[path]; ok {
		return apiPath
	}
	return ApiPathUnKnow
}
