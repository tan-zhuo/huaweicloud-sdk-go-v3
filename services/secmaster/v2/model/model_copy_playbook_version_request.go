package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyPlaybookVersionRequest Request Object
type CopyPlaybookVersionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	Body *CopyPlaybookInfo `json:"body,omitempty"`
}

func (o CopyPlaybookVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyPlaybookVersionRequest struct{}"
	}

	return strings.Join([]string{"CopyPlaybookVersionRequest", string(data)}, " ")
}
