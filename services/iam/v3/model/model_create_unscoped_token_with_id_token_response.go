package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUnscopedTokenWithIdTokenResponse Response Object
type CreateUnscopedTokenWithIdTokenResponse struct {
	Token *UnscopedTokenInfo `json:"token,omitempty"`

	XSubjectToken  *string `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUnscopedTokenWithIdTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUnscopedTokenWithIdTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateUnscopedTokenWithIdTokenResponse", string(data)}, " ")
}
