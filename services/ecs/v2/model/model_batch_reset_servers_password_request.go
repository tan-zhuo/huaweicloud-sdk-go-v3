package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResetServersPasswordRequest Request Object
type BatchResetServersPasswordRequest struct {
	Body *BatchResetServersPasswordRequestBody `json:"body,omitempty"`
}

func (o BatchResetServersPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetServersPasswordRequest struct{}"
	}

	return strings.Join([]string{"BatchResetServersPasswordRequest", string(data)}, " ")
}
