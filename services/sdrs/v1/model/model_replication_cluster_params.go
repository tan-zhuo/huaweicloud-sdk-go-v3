package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplicationClusterParams 复制集群相关参数
type ReplicationClusterParams struct {

	// 可用区名称。
	AvailabilityZone string `json:"availability_zone"`
}

func (o ReplicationClusterParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationClusterParams struct{}"
	}

	return strings.Join([]string{"ReplicationClusterParams", string(data)}, " ")
}
