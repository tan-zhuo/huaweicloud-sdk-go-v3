package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTranscodingsTemplateRequest Request Object
type UpdateTranscodingsTemplateRequest struct {
	Body *StreamTranscodingTemplate `json:"body,omitempty"`
}

func (o UpdateTranscodingsTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTranscodingsTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTranscodingsTemplateRequest", string(data)}, " ")
}
