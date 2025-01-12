package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrDeleteDeviceInGroupResponse Response Object
type CreateOrDeleteDeviceInGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOrDeleteDeviceInGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrDeleteDeviceInGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateOrDeleteDeviceInGroupResponse", string(data)}, " ")
}
