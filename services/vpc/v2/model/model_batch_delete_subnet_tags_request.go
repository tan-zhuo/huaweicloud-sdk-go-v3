package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubnetTagsRequest Request Object
type BatchDeleteSubnetTagsRequest struct {

	// 子网ID
	SubnetId string `json:"subnet_id"`

	Body *BatchDeleteSubnetTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubnetTagsRequest", string(data)}, " ")
}
