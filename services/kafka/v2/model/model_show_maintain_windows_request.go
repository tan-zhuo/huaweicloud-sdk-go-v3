package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMaintainWindowsRequest Request Object
type ShowMaintainWindowsRequest struct {
}

func (o ShowMaintainWindowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMaintainWindowsRequest struct{}"
	}

	return strings.Join([]string{"ShowMaintainWindowsRequest", string(data)}, " ")
}
