package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceNameRequest Request Object
type UpdateInstanceNameRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	Body *ModifyInstanceNameReq `json:"body,omitempty"`
}

func (o UpdateInstanceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequest", string(data)}, " ")
}
