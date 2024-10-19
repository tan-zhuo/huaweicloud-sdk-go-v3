package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaResponse Response Object
type ShowQuotaResponse struct {
	Quotas         *QuotaParams `json:"quotas,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponse", string(data)}, " ")
}
