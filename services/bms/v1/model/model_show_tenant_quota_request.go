package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantQuotaRequest Request Object
type ShowTenantQuotaRequest struct {
}

func (o ShowTenantQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantQuotaRequest", string(data)}, " ")
}
