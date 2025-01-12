package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RejectVpcPeeringRequest Request Object
type RejectVpcPeeringRequest struct {

	// 对等连接ID
	PeeringId string `json:"peering_id"`
}

func (o RejectVpcPeeringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectVpcPeeringRequest struct{}"
	}

	return strings.Join([]string{"RejectVpcPeeringRequest", string(data)}, " ")
}
