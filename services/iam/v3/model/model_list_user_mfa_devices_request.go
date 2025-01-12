package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserMfaDevicesRequest Request Object
type ListUserMfaDevicesRequest struct {
}

func (o ListUserMfaDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserMfaDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListUserMfaDevicesRequest", string(data)}, " ")
}
