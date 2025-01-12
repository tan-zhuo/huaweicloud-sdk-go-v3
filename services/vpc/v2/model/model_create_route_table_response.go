package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRouteTableResponse Response Object
type CreateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"CreateRouteTableResponse", string(data)}, " ")
}
