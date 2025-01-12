package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetIdTokenAuthParams auth信息
type GetIdTokenAuthParams struct {
	IdToken *GetIdTokenIdTokenBody `json:"id_token"`

	Scope *GetIdTokenIdScopeBody `json:"scope,omitempty"`
}

func (o GetIdTokenAuthParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetIdTokenAuthParams struct{}"
	}

	return strings.Join([]string{"GetIdTokenAuthParams", string(data)}, " ")
}
