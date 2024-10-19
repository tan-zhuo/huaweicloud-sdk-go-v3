package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPriceRequest Request Object
type ShowPriceRequest struct {
}

func (o ShowPriceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPriceRequest struct{}"
	}

	return strings.Join([]string{"ShowPriceRequest", string(data)}, " ")
}
