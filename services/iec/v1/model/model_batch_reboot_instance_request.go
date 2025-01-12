package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebootInstanceRequest Request Object
type BatchRebootInstanceRequest struct {
	Body *BatchRebootInstanceRequestBody `json:"body,omitempty"`
}

func (o BatchRebootInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootInstanceRequest", string(data)}, " ")
}
