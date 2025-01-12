package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoomSettingRequest Request Object
type UpdateRoomSettingRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 网络研讨会ID。
	ConferenceId string `json:"conference_id"`

	Body *OpenRoomSettingReq `json:"body,omitempty"`
}

func (o UpdateRoomSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoomSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateRoomSettingRequest", string(data)}, " ")
}
