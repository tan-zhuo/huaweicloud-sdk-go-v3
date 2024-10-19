package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInternetBandwidthRequestBody struct {
	InternetBandwidth *UpdateInternetBandwidthRequestBodyInternetBandwidth `json:"internet_bandwidth"`
}

func (o UpdateInternetBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternetBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInternetBandwidthRequestBody", string(data)}, " ")
}
