package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartFailoverProtectionGroupRequest Request Object
type StartFailoverProtectionGroupRequest struct {

	// 保护组的ID。
	ServerGroupId string `json:"server_group_id"`

	Body *FailoverProtectionGroupRequestBody `json:"body,omitempty"`
}

func (o StartFailoverProtectionGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartFailoverProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"StartFailoverProtectionGroupRequest", string(data)}, " ")
}
