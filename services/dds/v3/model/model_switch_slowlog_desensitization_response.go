package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSlowlogDesensitizationResponse Response Object
type SwitchSlowlogDesensitizationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchSlowlogDesensitizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSlowlogDesensitizationResponse struct{}"
	}

	return strings.Join([]string{"SwitchSlowlogDesensitizationResponse", string(data)}, " ")
}
