package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasRequest Request Object
type ShowQuotasRequest struct {
}

func (o ShowQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}
