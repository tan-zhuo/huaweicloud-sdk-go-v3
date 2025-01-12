package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartClusterRequest Request Object
type RestartClusterRequest struct {

	// 指定待重启集群的ID
	ClusterId string `json:"cluster_id"`

	Body *RestartClusterRequestBody `json:"body,omitempty"`
}

func (o RestartClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartClusterRequest struct{}"
	}

	return strings.Join([]string{"RestartClusterRequest", string(data)}, " ")
}
