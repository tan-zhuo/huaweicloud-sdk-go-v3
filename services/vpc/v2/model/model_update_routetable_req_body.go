package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoutetableReqBody
type UpdateRoutetableReqBody struct {
	Routetable *UpdateRouteTableReq `json:"routetable"`
}

func (o UpdateRoutetableReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutetableReqBody struct{}"
	}

	return strings.Join([]string{"UpdateRoutetableReqBody", string(data)}, " ")
}
