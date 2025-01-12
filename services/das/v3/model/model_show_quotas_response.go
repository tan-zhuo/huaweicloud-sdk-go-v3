package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasResponse Response Object
type ShowQuotasResponse struct {

	// 配额列表对象。
	Quotas         *interface{} `json:"quotas,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
