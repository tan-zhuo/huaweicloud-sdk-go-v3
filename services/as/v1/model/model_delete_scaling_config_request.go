package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScalingConfigRequest Request Object
type DeleteScalingConfigRequest struct {

	// 伸缩配置ID。
	ScalingConfigurationId string `json:"scaling_configuration_id"`
}

func (o DeleteScalingConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingConfigRequest", string(data)}, " ")
}
