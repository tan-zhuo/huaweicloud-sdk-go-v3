package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotPolicyRequest Request Object
type CreateSnapshotPolicyRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *BackupPolicy `json:"body,omitempty"`
}

func (o CreateSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateSnapshotPolicyRequest", string(data)}, " ")
}
