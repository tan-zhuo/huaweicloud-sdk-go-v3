package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWatermarkTemplateResponse Response Object
type DeleteWatermarkTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWatermarkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteWatermarkTemplateResponse", string(data)}, " ")
}
