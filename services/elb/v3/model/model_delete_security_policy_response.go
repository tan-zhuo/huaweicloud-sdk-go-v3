package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityPolicyResponse Response Object
type DeleteSecurityPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityPolicyResponse", string(data)}, " ")
}
