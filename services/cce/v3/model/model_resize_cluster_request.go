package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterRequest Request Object
type ResizeClusterRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// 集群状态兼容Error参数，用于API平滑切换。 兼容场景下，errorStatus为空则屏蔽Error状态为Deleting状态。
	ErrorStatus *string `json:"errorStatus,omitempty"`

	Body *ResizeClusterRequestBody `json:"body,omitempty"`
}

func (o ResizeClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterRequest struct{}"
	}

	return strings.Join([]string{"ResizeClusterRequest", string(data)}, " ")
}
