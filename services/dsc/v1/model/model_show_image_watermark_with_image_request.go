package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageWatermarkWithImageRequest Request Object
type ShowImageWatermarkWithImageRequest struct {
	Body *ShowImageWatermarkWithImageRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ShowImageWatermarkWithImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkWithImageRequest struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkWithImageRequest", string(data)}, " ")
}
