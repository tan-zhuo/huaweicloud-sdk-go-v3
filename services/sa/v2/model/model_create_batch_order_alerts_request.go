package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBatchOrderAlertsRequest Request Object
type CreateBatchOrderAlertsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	Body *OrderAlert `json:"body,omitempty"`
}

func (o CreateBatchOrderAlertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchOrderAlertsRequest struct{}"
	}

	return strings.Join([]string{"CreateBatchOrderAlertsRequest", string(data)}, " ")
}
