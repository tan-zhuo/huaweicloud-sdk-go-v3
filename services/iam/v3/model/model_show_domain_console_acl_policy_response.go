package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainConsoleAclPolicyResponse Response Object
type ShowDomainConsoleAclPolicyResponse struct {
	ConsoleAclPolicy *AclPolicyResult `json:"console_acl_policy,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o ShowDomainConsoleAclPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainConsoleAclPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainConsoleAclPolicyResponse", string(data)}, " ")
}
