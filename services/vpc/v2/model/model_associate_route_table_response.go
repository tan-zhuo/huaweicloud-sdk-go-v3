package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRouteTableResponse Response Object
type AssociateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o AssociateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"AssociateRouteTableResponse", string(data)}, " ")
}
