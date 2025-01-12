package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScopedTokenAuth
type ScopedTokenAuth struct {
	Identity *ScopedTokenIdentity `json:"identity"`

	Scope *TokenSocpeOption `json:"scope"`
}

func (o ScopedTokenAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScopedTokenAuth struct{}"
	}

	return strings.Join([]string{"ScopedTokenAuth", string(data)}, " ")
}
