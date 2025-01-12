package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateServerVirtualIpResponse Response Object
type DisassociateServerVirtualIpResponse struct {

	// 云服务器网卡ID
	PortId         *string `json:"port_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateServerVirtualIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateServerVirtualIpResponse struct{}"
	}

	return strings.Join([]string{"DisassociateServerVirtualIpResponse", string(data)}, " ")
}
