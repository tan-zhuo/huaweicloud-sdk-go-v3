package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSharedBandwidthsRequest Request Object
type BatchCreateSharedBandwidthsRequest struct {
	Body *BatchCreateBandwidthRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSharedBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSharedBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSharedBandwidthsRequest", string(data)}, " ")
}
