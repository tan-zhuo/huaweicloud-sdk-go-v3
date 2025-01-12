package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateOrDeleteTagsRequest Request Object
type BatchCreateOrDeleteTagsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CreateOrDeleteInstanceTags `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteTagsRequest", string(data)}, " ")
}
