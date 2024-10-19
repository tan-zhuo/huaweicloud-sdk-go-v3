package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePersonalAuthResponse Response Object
type CreatePersonalAuthResponse struct {
	Authorization  *AuthorizationVi `json:"authorization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreatePersonalAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePersonalAuthResponse struct{}"
	}

	return strings.Join([]string{"CreatePersonalAuthResponse", string(data)}, " ")
}
