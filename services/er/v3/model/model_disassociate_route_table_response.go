package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateRouteTableResponse Response Object
type DisassociateRouteTableResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"DisassociateRouteTableResponse", string(data)}, " ")
}
