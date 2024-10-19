package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteClouddcnSubnetsTagsRequest Request Object
type BatchDeleteClouddcnSubnetsTagsRequest struct {

	// Clouddcn子网的id
	ResourceId string `json:"resource_id"`

	Body *BatchDeleteRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteClouddcnSubnetsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClouddcnSubnetsTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteClouddcnSubnetsTagsRequest", string(data)}, " ")
}
