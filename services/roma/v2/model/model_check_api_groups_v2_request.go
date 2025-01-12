package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckApiGroupsV2Request Request Object
type CheckApiGroupsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ApiGroupCheck `json:"body,omitempty"`
}

func (o CheckApiGroupsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckApiGroupsV2Request struct{}"
	}

	return strings.Join([]string{"CheckApiGroupsV2Request", string(data)}, " ")
}
