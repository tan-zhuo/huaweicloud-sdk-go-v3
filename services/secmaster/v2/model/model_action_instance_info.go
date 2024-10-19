package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionInstanceInfo 流程实例
type ActionInstanceInfo struct {
	Action *ActionInfo `json:"action,omitempty"`

	InstanceLog *AuditLogInfo `json:"instance_log,omitempty"`
}

func (o ActionInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionInstanceInfo struct{}"
	}

	return strings.Join([]string{"ActionInstanceInfo", string(data)}, " ")
}
