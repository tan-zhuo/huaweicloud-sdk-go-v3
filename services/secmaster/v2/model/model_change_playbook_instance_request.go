package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePlaybookInstanceRequest Request Object
type ChangePlaybookInstanceRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本实例ID
	InstanceId string `json:"instance_id"`

	Body *OperationPlaybookInfo `json:"body,omitempty"`
}

func (o ChangePlaybookInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePlaybookInstanceRequest struct{}"
	}

	return strings.Join([]string{"ChangePlaybookInstanceRequest", string(data)}, " ")
}
