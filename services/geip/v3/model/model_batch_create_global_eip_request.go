package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGlobalEipRequest Request Object
type BatchCreateGlobalEipRequest struct {
	Body *BatchCreateGlobalEipRequestBody `json:"body,omitempty"`
}

func (o BatchCreateGlobalEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEipRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEipRequest", string(data)}, " ")
}
