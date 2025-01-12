package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsscoiateReq
type AsscoiateReq struct {
	Subnets *AssociateRouteTableAndSubnetReq `json:"subnets"`
}

func (o AsscoiateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsscoiateReq struct{}"
	}

	return strings.Join([]string{"AsscoiateReq", string(data)}, " ")
}
