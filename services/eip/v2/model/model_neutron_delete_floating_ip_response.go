package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteFloatingIpResponse Response Object
type NeutronDeleteFloatingIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFloatingIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFloatingIpResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFloatingIpResponse", string(data)}, " ")
}
