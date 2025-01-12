package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWatermarkTemplateResponse Response Object
type UpdateWatermarkTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWatermarkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateWatermarkTemplateResponse", string(data)}, " ")
}
