package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityGroupResponse Response Object
type UpdateSecurityGroupResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	SecurityGroup  *SecurityGroupInfo `json:"security_group,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupResponse", string(data)}, " ")
}
