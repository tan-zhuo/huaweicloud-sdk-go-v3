package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteProtectedInstancesRequest Request Object
type BatchDeleteProtectedInstancesRequest struct {
	Body *BatchDeleteProtectedInstancesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteProtectedInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedInstancesRequest", string(data)}, " ")
}
