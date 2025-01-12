package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevicesResponse Response Object
type ListDevicesResponse struct {

	// 满足条件的设备总数
	Total *int32 `json:"total,omitempty"`

	Data           *[]GetDevicesListArrayObject `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListDevicesResponse", string(data)}, " ")
}
