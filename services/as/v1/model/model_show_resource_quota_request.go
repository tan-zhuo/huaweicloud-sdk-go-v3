package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceQuotaRequest Request Object
type ShowResourceQuotaRequest struct {
}

func (o ShowResourceQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceQuotaRequest", string(data)}, " ")
}
