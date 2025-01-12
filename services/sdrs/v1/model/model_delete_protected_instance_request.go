package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectedInstanceRequest Request Object
type DeleteProtectedInstanceRequest struct {

	// 保护实例的ID。
	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *DeleteProtectedInstanceRequestBody `json:"body,omitempty"`
}

func (o DeleteProtectedInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceRequest", string(data)}, " ")
}
