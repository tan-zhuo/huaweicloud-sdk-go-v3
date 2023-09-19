package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nic 租户自定义的网卡的结构体，为待创建的云服务器的网卡信息。
type Nic struct {

	// 租户自定义的子网 ID，为待创建的云服务器所属的子网。  需要指定tenant_vpc_id对应VPC下已创建的子网（subnet）的网络ID，UUID格式。
	SubnetId string `json:"subnet_id"`

	// 是否支持ipv6。  取值为true时，标识此网卡支持ipv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	Ipv6Bandwidth *NicIpv6Bandwidth `json:"ipv6_bandwidth,omitempty"`
}

func (o Nic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nic struct{}"
	}

	return strings.Join([]string{"Nic", string(data)}, " ")
}
