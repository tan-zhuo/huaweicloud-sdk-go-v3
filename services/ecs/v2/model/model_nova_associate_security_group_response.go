package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaAssociateSecurityGroupResponse Response Object
type NovaAssociateSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaAssociateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAssociateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NovaAssociateSecurityGroupResponse", string(data)}, " ")
}
