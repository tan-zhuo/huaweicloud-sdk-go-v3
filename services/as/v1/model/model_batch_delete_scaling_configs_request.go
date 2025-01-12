package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScalingConfigsRequest Request Object
type BatchDeleteScalingConfigsRequest struct {
	Body *BatchDeleteScalingConfigOption `json:"body,omitempty"`
}

func (o BatchDeleteScalingConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingConfigsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingConfigsRequest", string(data)}, " ")
}
