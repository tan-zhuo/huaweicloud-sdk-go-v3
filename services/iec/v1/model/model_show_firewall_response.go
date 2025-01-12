package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFirewallResponse Response Object
type ShowFirewallResponse struct {
	Firewall       *Firewall `json:"firewall,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFirewallResponse struct{}"
	}

	return strings.Join([]string{"ShowFirewallResponse", string(data)}, " ")
}
