package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnUserGroupResponse Response Object
type DeleteVpnUserGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpnUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnUserGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpnUserGroupResponse", string(data)}, " ")
}
