package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWatermarkTemplateRequest Request Object
type CreateWatermarkTemplateRequest struct {
	Body *WatermarkTemplate `json:"body,omitempty"`
}

func (o CreateWatermarkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateWatermarkTemplateRequest", string(data)}, " ")
}
