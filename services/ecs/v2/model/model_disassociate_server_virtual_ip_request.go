package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisassociateServerVirtualIpRequest struct {
	// 云服务器网卡ID。

	NicId string `json:"nic_id"`

	Body *DisassociateServerVirtualIpRequestBody `json:"body,omitempty"`
}

func (o DisassociateServerVirtualIpRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateServerVirtualIpRequest struct{}"
	}

	return strings.Join([]string{"DisassociateServerVirtualIpRequest", string(data)}, " ")
}
