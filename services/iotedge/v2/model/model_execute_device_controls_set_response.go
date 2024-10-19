package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDeviceControlsSetResponse Response Object
type ExecuteDeviceControlsSetResponse struct {

	// 属性设置的响应码，具体为实际设备返回的响应码
	ResultCode *int32 `json:"result_code,omitempty"`

	// 属性设置的描述，具体为实际设备返回的描述
	ResultDesc     *string `json:"result_desc,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteDeviceControlsSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeviceControlsSetResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDeviceControlsSetResponse", string(data)}, " ")
}
