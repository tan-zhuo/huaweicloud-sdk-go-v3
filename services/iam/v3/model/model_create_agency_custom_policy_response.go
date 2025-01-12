package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyCustomPolicyResponse Response Object
type CreateAgencyCustomPolicyResponse struct {
	Role           *AgencyPolicyRoleResult `json:"role,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreateAgencyCustomPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateAgencyCustomPolicyResponse", string(data)}, " ")
}
