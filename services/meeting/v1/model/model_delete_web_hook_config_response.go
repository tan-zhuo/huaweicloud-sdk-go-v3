package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWebHookConfigResponse Response Object
type DeleteWebHookConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWebHookConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWebHookConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteWebHookConfigResponse", string(data)}, " ")
}
