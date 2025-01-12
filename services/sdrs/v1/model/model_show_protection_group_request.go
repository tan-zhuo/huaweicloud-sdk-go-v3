package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProtectionGroupRequest Request Object
type ShowProtectionGroupRequest struct {

	// 保护组的ID。
	ServerGroupId string `json:"server_group_id"`
}

func (o ShowProtectionGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowProtectionGroupRequest", string(data)}, " ")
}
