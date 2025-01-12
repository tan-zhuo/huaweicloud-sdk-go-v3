package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUnprotectScalingInstancesResponse Response Object
type BatchUnprotectScalingInstancesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUnprotectScalingInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnprotectScalingInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchUnprotectScalingInstancesResponse", string(data)}, " ")
}
