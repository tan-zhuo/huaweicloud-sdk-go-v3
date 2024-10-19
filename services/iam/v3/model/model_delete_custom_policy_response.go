package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomPolicyResponse Response Object
type DeleteCustomPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomPolicyResponse", string(data)}, " ")
}
