package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlQuotasRequest Request Object
type UpdateGaussMySqlQuotasRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *SetQuotasRequestBody `json:"body,omitempty"`
}

func (o UpdateGaussMySqlQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlQuotasRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlQuotasRequest", string(data)}, " ")
}
