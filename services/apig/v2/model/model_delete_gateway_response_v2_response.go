package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGatewayResponseV2Response Response Object
type DeleteGatewayResponseV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGatewayResponseV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGatewayResponseV2Response struct{}"
	}

	return strings.Join([]string{"DeleteGatewayResponseV2Response", string(data)}, " ")
}
