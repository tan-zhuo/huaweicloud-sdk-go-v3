package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRouteTableRequest Request Object
type DeleteRouteTableRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id"`
}

func (o DeleteRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRouteTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteRouteTableRequest", string(data)}, " ")
}
