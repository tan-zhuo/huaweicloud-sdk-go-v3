package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAutopilotClusterTagsRequest Request Object
type BatchDeleteAutopilotClusterTagsRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	Body *BatchDeleteClusterTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteAutopilotClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAutopilotClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAutopilotClusterTagsRequest", string(data)}, " ")
}
