package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceConfigurationRequest Request Object
type ShowInstanceConfigurationRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationRequest", string(data)}, " ")
}
