package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectionGroupRequest Request Object
type DeleteProtectionGroupRequest struct {

	// 保护组ID。
	ServerGroupId string `json:"server_group_id"`
}

func (o DeleteProtectionGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtectionGroupRequest", string(data)}, " ")
}
