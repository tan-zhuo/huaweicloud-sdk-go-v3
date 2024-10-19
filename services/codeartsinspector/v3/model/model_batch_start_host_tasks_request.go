package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartHostTasksRequest Request Object
type BatchStartHostTasksRequest struct {
	Body *BatchStartHostTasksRequestBody `json:"body,omitempty"`
}

func (o BatchStartHostTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartHostTasksRequest struct{}"
	}

	return strings.Join([]string{"BatchStartHostTasksRequest", string(data)}, " ")
}
