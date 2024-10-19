package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSparkJobTemplateRequest Request Object
type UpdateSparkJobTemplateRequest struct {
	TemplateId string `json:"template_id"`

	Body *UpdateSparkJobTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateSparkJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSparkJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateSparkJobTemplateRequest", string(data)}, " ")
}
