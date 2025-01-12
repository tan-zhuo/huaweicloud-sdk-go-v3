package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainApiAclPolicyResponse Response Object
type ShowDomainApiAclPolicyResponse struct {
	ApiAclPolicy   *AclPolicyResult `json:"api_acl_policy,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDomainApiAclPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainApiAclPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainApiAclPolicyResponse", string(data)}, " ")
}
