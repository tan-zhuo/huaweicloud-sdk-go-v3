package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVirtualGatewayRequestBody struct {
	VirtualGateway *CreateVirtualGateway `json:"virtual_gateway,omitempty"`
}

func (o CreateVirtualGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVirtualGatewayRequestBody", string(data)}, " ")
}
