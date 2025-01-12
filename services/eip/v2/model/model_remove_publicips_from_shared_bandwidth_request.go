package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemovePublicipsFromSharedBandwidthRequest Request Object
type RemovePublicipsFromSharedBandwidthRequest struct {

	// 带宽唯一标识
	BandwidthId string `json:"bandwidth_id"`

	Body *RemovePublicipsFromSharedBandwidthRequestBody `json:"body,omitempty"`
}

func (o RemovePublicipsFromSharedBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemovePublicipsFromSharedBandwidthRequest struct{}"
	}

	return strings.Join([]string{"RemovePublicipsFromSharedBandwidthRequest", string(data)}, " ")
}
