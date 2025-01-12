package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBigkeyScanTaskRequest Request Object
type DeleteBigkeyScanTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 大key分析任务ID。
	BigkeyId string `json:"bigkey_id"`
}

func (o DeleteBigkeyScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBigkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteBigkeyScanTaskRequest", string(data)}, " ")
}
