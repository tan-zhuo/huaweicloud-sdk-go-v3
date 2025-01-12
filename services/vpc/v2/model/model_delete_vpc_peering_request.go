package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcPeeringRequest Request Object
type DeleteVpcPeeringRequest struct {

	// 对等连接ID
	PeeringId string `json:"peering_id"`
}

func (o DeleteVpcPeeringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcPeeringRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcPeeringRequest", string(data)}, " ")
}
