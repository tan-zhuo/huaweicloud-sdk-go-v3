package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateSecurityGroupRequest Request Object
type NeutronCreateSecurityGroupRequest struct {
	Body *NeutronCreateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRequest", string(data)}, " ")
}
