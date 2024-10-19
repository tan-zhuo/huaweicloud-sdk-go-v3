package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkConfigurationRequest Request Object
type ShowNetworkConfigurationRequest struct {
	Body *NetworkRequestBody `json:"body,omitempty"`
}

func (o ShowNetworkConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowNetworkConfigurationRequest", string(data)}, " ")
}
