package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyEpsQuotasRequest Request Object
type ModifyEpsQuotasRequest struct {
	Body *NoSqlModiflyEpsQuotasRequestBody `json:"body,omitempty"`
}

func (o ModifyEpsQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEpsQuotasRequest struct{}"
	}

	return strings.Join([]string{"ModifyEpsQuotasRequest", string(data)}, " ")
}
