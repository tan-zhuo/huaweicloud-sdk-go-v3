package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookActionRequest Request Object
type CreatePlaybookActionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	// Create actions
	Body *[]CreateAction `json:"body,omitempty"`
}

func (o CreatePlaybookActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookActionRequest struct{}"
	}

	return strings.Join([]string{"CreatePlaybookActionRequest", string(data)}, " ")
}
