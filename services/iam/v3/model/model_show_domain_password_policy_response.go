package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainPasswordPolicyResponse Response Object
type ShowDomainPasswordPolicyResponse struct {
	PasswordPolicy *PasswordPolicyResult `json:"password_policy,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowDomainPasswordPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainPasswordPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainPasswordPolicyResponse", string(data)}, " ")
}
