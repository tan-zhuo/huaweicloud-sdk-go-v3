package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartManagerRequest Request Object
type RestartManagerRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o RestartManagerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartManagerRequest struct{}"
	}

	return strings.Join([]string{"RestartManagerRequest", string(data)}, " ")
}
