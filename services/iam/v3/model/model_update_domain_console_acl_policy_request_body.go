package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainConsoleAclPolicyRequestBody
type UpdateDomainConsoleAclPolicyRequestBody struct {
	ConsoleAclPolicy *AclPolicyOption `json:"console_acl_policy"`
}

func (o UpdateDomainConsoleAclPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainConsoleAclPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainConsoleAclPolicyRequestBody", string(data)}, " ")
}
