package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGaussMySqlProjectQuotasResponse Response Object
type ShowGaussMySqlProjectQuotasResponse struct {
	Quotas         *ProjectQuotas `json:"quotas,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowGaussMySqlProjectQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlProjectQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlProjectQuotasResponse", string(data)}, " ")
}
