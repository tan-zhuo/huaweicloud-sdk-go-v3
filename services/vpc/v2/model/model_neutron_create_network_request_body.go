package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateNetworkRequestBody
type NeutronCreateNetworkRequestBody struct {
	Network *NeutronCreateNetworkOption `json:"network"`
}

func (o NeutronCreateNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateNetworkRequestBody", string(data)}, " ")
}
