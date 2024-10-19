package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchInstancePrimaryResponse Response Object
type SwitchInstancePrimaryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchInstancePrimaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchInstancePrimaryResponse struct{}"
	}

	return strings.Join([]string{"SwitchInstancePrimaryResponse", string(data)}, " ")
}
