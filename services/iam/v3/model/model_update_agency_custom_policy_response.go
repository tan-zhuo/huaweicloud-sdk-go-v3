package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyCustomPolicyResponse Response Object
type UpdateAgencyCustomPolicyResponse struct {
	Role           *AgencyPolicyRoleResult `json:"role,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateAgencyCustomPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgencyCustomPolicyResponse", string(data)}, " ")
}
