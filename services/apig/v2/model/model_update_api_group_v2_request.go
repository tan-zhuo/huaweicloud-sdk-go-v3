package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApiGroupV2Request Request Object
type UpdateApiGroupV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	Body *ApiGroupBase `json:"body,omitempty"`
}

func (o UpdateApiGroupV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"UpdateApiGroupV2Request", string(data)}, " ")
}
