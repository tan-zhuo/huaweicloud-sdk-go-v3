package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBandwidthRequest Request Object
type ShowBandwidthRequest struct {

	// 带宽ID。
	BandwidthId string `json:"bandwidth_id"`
}

func (o ShowBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthRequest struct{}"
	}

	return strings.Join([]string{"ShowBandwidthRequest", string(data)}, " ")
}
