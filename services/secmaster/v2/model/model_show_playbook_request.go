package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookRequest Request Object
type ShowPlaybookRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// ID of playbook
	PlaybookId string `json:"playbook_id"`
}

func (o ShowPlaybookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookRequest", string(data)}, " ")
}
