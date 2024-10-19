package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetInstancePasswordRequest Request Object
type ResetInstancePasswordRequest struct {
	Body *ResetPassword `json:"body,omitempty"`
}

func (o ResetInstancePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetInstancePasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetInstancePasswordRequest", string(data)}, " ")
}
