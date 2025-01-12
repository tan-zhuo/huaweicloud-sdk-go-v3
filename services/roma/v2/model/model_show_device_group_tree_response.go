package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceGroupTreeResponse Response Object
type ShowDeviceGroupTreeResponse struct {

	// 本次返回数量
	Size *int32 `json:"size,omitempty"`

	// 设备分组信息
	Items          *[]GroupTreeResponse `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowDeviceGroupTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceGroupTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceGroupTreeResponse", string(data)}, " ")
}
