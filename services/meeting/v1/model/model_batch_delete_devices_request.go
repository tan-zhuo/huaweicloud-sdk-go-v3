package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDevicesRequest Request Object
type BatchDeleteDevicesRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchDeleteDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDevicesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDevicesRequest", string(data)}, " ")
}
