package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachInternetBandwidthRequest Request Object
type BatchAttachInternetBandwidthRequest struct {
	Body *BatchAttachInternetBandwidthsGlobalEipRequestBody `json:"body,omitempty"`
}

func (o BatchAttachInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"BatchAttachInternetBandwidthRequest", string(data)}, " ")
}
