package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookVersionRequest Request Object
type UpdatePlaybookVersionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	Body *ModifyPlaybookVersionInfo `json:"body,omitempty"`
}

func (o UpdatePlaybookVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookVersionRequest struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookVersionRequest", string(data)}, " ")
}
