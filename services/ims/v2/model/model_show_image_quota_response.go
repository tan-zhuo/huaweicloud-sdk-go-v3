package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageQuotaResponse Response Object
type ShowImageQuotaResponse struct {
	Quotas         *Quota `json:"quotas,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowImageQuotaResponse", string(data)}, " ")
}
