package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProtectedInstanceRequest Request Object
type ShowProtectedInstanceRequest struct {

	// 保护实例的ID。
	ProtectedInstanceId string `json:"protected_instance_id"`
}

func (o ShowProtectedInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectedInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowProtectedInstanceRequest", string(data)}, " ")
}
