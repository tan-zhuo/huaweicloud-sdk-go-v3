package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyTemplate struct {
	IkePolicy *VgwIkePolicy `json:"ike_policy,omitempty"`

	IpsecPolicy *VgwIpsecPolicy `json:"ipsec_policy,omitempty"`
}

func (o PolicyTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTemplate struct{}"
	}

	return strings.Join([]string{"PolicyTemplate", string(data)}, " ")
}
