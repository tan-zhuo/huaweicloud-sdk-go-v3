package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunQueueActionRequest Request Object
type RunQueueActionRequest struct {
	QueueName string `json:"queue_name"`

	Body *RunQueueActionRequestBody `json:"body,omitempty"`
}

func (o RunQueueActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueueActionRequest struct{}"
	}

	return strings.Join([]string{"RunQueueActionRequest", string(data)}, " ")
}
