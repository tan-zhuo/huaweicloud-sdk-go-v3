package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePoolRequest Request Object
type CreatePoolRequest struct {
	Body *CreatePoolRequestBody `json:"body,omitempty"`
}

func (o CreatePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolRequest struct{}"
	}

	return strings.Join([]string{"CreatePoolRequest", string(data)}, " ")
}
