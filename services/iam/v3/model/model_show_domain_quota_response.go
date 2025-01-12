package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainQuotaResponse Response Object
type ShowDomainQuotaResponse struct {
	Quotas         *QuotaResult `json:"quotas,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDomainQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaResponse", string(data)}, " ")
}
