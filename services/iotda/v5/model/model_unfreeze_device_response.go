package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnfreezeDeviceResponse Response Object
type UnfreezeDeviceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnfreezeDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeDeviceResponse struct{}"
	}

	return strings.Join([]string{"UnfreezeDeviceResponse", string(data)}, " ")
}
