package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClientNetworkRequestBody struct {

	// 客户端所在网段。  - 跨网段访问配置只有在客户端与副本集实例部署在不同网段的情况下才需要配置，例如访问副本集的客户端所在网段为192.168.0.0/16，副本集所在的网段为172.16.0.0/24，则需要添加跨网段配置192.168.0.0/16才能正常访问。  - 例如配置的源端网段为192.168.0.0/xx，则xx的输入值必须在8到32之间。  - 源端ECS连接实例的前提是与实例节点网络通信正常，如果网络不通，可以参考对等连接进行相关配置。
	ClientNetworkRanges []string `json:"client_network_ranges"`
}

func (o UpdateClientNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateClientNetworkRequestBody", string(data)}, " ")
}
