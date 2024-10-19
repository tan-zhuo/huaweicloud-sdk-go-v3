package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWatermarkTemplateRequest Request Object
type DeleteWatermarkTemplateRequest struct {

	// 水印模板ID
	TemplateId int32 `json:"template_id"`
}

func (o DeleteWatermarkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteWatermarkTemplateRequest", string(data)}, " ")
}
