package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScalingConfigResponse Response Object
type ShowScalingConfigResponse struct {
	ScalingConfiguration *ScalingConfiguration `json:"scaling_configuration,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowScalingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowScalingConfigResponse", string(data)}, " ")
}
