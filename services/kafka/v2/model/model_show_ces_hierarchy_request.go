package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCesHierarchyRequest Request Object
type ShowCesHierarchyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowCesHierarchyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCesHierarchyRequest struct{}"
	}

	return strings.Join([]string{"ShowCesHierarchyRequest", string(data)}, " ")
}
