package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesRequest Request Object
type ListInstancesRequest struct {

	// 云堡垒机实例ID。（非必传，需要查询单个实例详情时传入）
	InstanceId float32 `json:"instance_id,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
