package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationPolicyResponse Response Object
type ShowOrganizationPolicyResponse struct {
	Policy         *OrganizationPolicy `json:"policy,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowOrganizationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationPolicyResponse", string(data)}, " ")
}
