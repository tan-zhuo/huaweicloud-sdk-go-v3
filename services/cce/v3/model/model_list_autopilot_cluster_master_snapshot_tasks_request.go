package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutopilotClusterMasterSnapshotTasksRequest Request Object
type ListAutopilotClusterMasterSnapshotTasksRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`
}

func (o ListAutopilotClusterMasterSnapshotTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutopilotClusterMasterSnapshotTasksRequest struct{}"
	}

	return strings.Join([]string{"ListAutopilotClusterMasterSnapshotTasksRequest", string(data)}, " ")
}
