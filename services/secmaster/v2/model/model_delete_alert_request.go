package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRequest Request Object
type DeleteAlertRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *DeleteAlertRequestBody `json:"body,omitempty"`
}

func (o DeleteAlertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlertRequest", string(data)}, " ")
}
