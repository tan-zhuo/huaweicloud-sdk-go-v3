package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageWatermarkWithImageByAddressRequest Request Object
type ShowImageWatermarkWithImageByAddressRequest struct {
	Body *ShowImageWatermarkWithImageByAddressRequestBody `json:"body,omitempty"`
}

func (o ShowImageWatermarkWithImageByAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkWithImageByAddressRequest struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkWithImageByAddressRequest", string(data)}, " ")
}
