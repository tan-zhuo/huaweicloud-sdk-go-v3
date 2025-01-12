package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGatewayResponseTypeV2Response Response Object
type UpdateGatewayResponseTypeV2Response struct {
	Body           map[string]ResponseInfoResp `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateGatewayResponseTypeV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGatewayResponseTypeV2Response struct{}"
	}

	return strings.Join([]string{"UpdateGatewayResponseTypeV2Response", string(data)}, " ")
}
