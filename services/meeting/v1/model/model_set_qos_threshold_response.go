package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetQosThresholdResponse Response Object
type SetQosThresholdResponse struct {

	// 返回码
	Code *string `json:"code,omitempty"`

	// 返回信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetQosThresholdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetQosThresholdResponse struct{}"
	}

	return strings.Join([]string{"SetQosThresholdResponse", string(data)}, " ")
}
