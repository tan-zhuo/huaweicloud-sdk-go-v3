package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnConnectionRequest Request Object
type DeleteVpnConnectionRequest struct {

	// vpn连接ID
	VpnConnectionId string `json:"vpn_connection_id"`
}

func (o DeleteVpnConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpnConnectionRequest", string(data)}, " ")
}
