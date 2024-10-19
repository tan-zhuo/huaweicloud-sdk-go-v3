package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertsRequest Request Object
type ListAlertsRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *DataobjectSearch `json:"body,omitempty"`
}

func (o ListAlertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertsRequest struct{}"
	}

	return strings.Join([]string{"ListAlertsRequest", string(data)}, " ")
}
