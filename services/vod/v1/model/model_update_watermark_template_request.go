package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWatermarkTemplateRequest Request Object
type UpdateWatermarkTemplateRequest struct {
	Body *UpdateWatermarkTemplateReq `json:"body,omitempty"`
}

func (o UpdateWatermarkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateWatermarkTemplateRequest", string(data)}, " ")
}
