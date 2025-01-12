package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainApiAclPolicyResponse Response Object
type UpdateDomainApiAclPolicyResponse struct {
	ApiAclPolicy   *AclPolicyResult `json:"api_acl_policy,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateDomainApiAclPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainApiAclPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainApiAclPolicyResponse", string(data)}, " ")
}
