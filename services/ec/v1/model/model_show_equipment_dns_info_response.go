package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEquipmentDnsInfoResponse Response Object
type ShowEquipmentDnsInfoResponse struct {

	// 主DNS
	MasterDns *string `json:"master_dns,omitempty"`

	// 备DNS
	SlaveDns       *string `json:"slave_dns,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEquipmentDnsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentDnsInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEquipmentDnsInfoResponse", string(data)}, " ")
}
