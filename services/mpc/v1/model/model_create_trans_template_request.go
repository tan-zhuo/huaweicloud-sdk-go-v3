package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransTemplateRequest Request Object
type CreateTransTemplateRequest struct {
	Body *TransTemplate `json:"body,omitempty"`
}

func (o CreateTransTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTransTemplateRequest", string(data)}, " ")
}
