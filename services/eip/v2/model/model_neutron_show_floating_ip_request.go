package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowFloatingIpRequest Request Object
type NeutronShowFloatingIpRequest struct {

	// floatingip的ID
	FloatingipId string `json:"floatingip_id"`
}

func (o NeutronShowFloatingIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFloatingIpRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowFloatingIpRequest", string(data)}, " ")
}
