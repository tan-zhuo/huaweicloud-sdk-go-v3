package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeyAliasRequest Request Object
type UpdateKeyAliasRequest struct {
	Body *UpdateKeyAliasRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyAliasRequest", string(data)}, " ")
}
