package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetScalingInstancesStandbyResponse Response Object
type BatchSetScalingInstancesStandbyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchSetScalingInstancesStandbyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetScalingInstancesStandbyResponse struct{}"
	}

	return strings.Join([]string{"BatchSetScalingInstancesStandbyResponse", string(data)}, " ")
}
