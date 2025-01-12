package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateServersNameRequest Request Object
type BatchUpdateServersNameRequest struct {
	Body *BatchUpdateServersNameRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateServersNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateServersNameRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateServersNameRequest", string(data)}, " ")
}
