package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotUserChartsQuotasRequest Request Object
type ShowAutopilotUserChartsQuotasRequest struct {
}

func (o ShowAutopilotUserChartsQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotUserChartsQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowAutopilotUserChartsQuotasRequest", string(data)}, " ")
}
