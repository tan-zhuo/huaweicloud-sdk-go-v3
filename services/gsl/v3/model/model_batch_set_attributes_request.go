package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetAttributesRequest Request Object
type BatchSetAttributesRequest struct {
	Body *BatchSetAttributesReq `json:"body,omitempty"`
}

func (o BatchSetAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetAttributesRequest struct{}"
	}

	return strings.Join([]string{"BatchSetAttributesRequest", string(data)}, " ")
}
