package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemovePublicipsFromSharedBandwidthResponse Response Object
type RemovePublicipsFromSharedBandwidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemovePublicipsFromSharedBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemovePublicipsFromSharedBandwidthResponse struct{}"
	}

	return strings.Join([]string{"RemovePublicipsFromSharedBandwidthResponse", string(data)}, " ")
}
