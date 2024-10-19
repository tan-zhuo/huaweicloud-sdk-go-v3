package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserChartsQuotasRequest Request Object
type ShowUserChartsQuotasRequest struct {
}

func (o ShowUserChartsQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserChartsQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowUserChartsQuotasRequest", string(data)}, " ")
}
