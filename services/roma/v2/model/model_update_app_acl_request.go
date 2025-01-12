package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppAclRequest Request Object
type UpdateAppAclRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`

	Body *AppAclCreate `json:"body,omitempty"`
}

func (o UpdateAppAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppAclRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppAclRequest", string(data)}, " ")
}
