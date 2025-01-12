package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestMuteReqBody 全场静音请求。
type RestMuteReqBody struct {

	// - 0: 取消静音 - 1: 静音
	IsMute int32 `json:"isMute"`

	// 是否允许自己解除静音（仅静音时有效），默认为允许。 - 0: 不允许 - 1: 允许
	AllowUnmuteByOneself *int32 `json:"allowUnmuteByOneself,omitempty"`
}

func (o RestMuteReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestMuteReqBody struct{}"
	}

	return strings.Join([]string{"RestMuteReqBody", string(data)}, " ")
}
