package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScalingV2PolicyResponse Response Object
type ShowScalingV2PolicyResponse struct {
	ScalingPolicy  *ScalingV2PolicyDetail `json:"scaling_policy,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowScalingV2PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingV2PolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowScalingV2PolicyResponse", string(data)}, " ")
}
