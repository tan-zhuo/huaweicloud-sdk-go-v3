package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetIdTokenIdScopeBody scope信息
type GetIdTokenIdScopeBody struct {
	Domain *GetIdTokenScopeDomainOrProjectBody `json:"domain,omitempty"`

	Project *GetIdTokenScopeDomainOrProjectBody `json:"project,omitempty"`
}

func (o GetIdTokenIdScopeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetIdTokenIdScopeBody struct{}"
	}

	return strings.Join([]string{"GetIdTokenIdScopeBody", string(data)}, " ")
}
