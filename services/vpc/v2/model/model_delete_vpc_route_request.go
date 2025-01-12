package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcRouteRequest Request Object
type DeleteVpcRouteRequest struct {

	// 路由ID
	RouteId string `json:"route_id"`
}

func (o DeleteVpcRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcRouteRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcRouteRequest", string(data)}, " ")
}
