package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEngineQuotasResponse Response Object
type ShowEngineQuotasResponse struct {
	Quotas         *EngineQuotaV2Quotas `json:"quotas,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowEngineQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowEngineQuotasResponse", string(data)}, " ")
}
