package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTasksRequest Request Object
type CountTasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o CountTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTasksRequest struct{}"
	}

	return strings.Join([]string{"CountTasksRequest", string(data)}, " ")
}
