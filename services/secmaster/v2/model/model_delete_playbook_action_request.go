package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookActionRequest Request Object
type DeletePlaybookActionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	// 剧本动作ID
	ActionId string `json:"action_id"`
}

func (o DeletePlaybookActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookActionRequest struct{}"
	}

	return strings.Join([]string{"DeletePlaybookActionRequest", string(data)}, " ")
}
