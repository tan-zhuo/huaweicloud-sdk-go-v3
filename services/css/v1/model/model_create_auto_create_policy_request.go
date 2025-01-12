package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutoCreatePolicyRequest Request Object
type CreateAutoCreatePolicyRequest struct {

	// 指定要自动创建快照的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *SetRdsBackupCnfReq `json:"body,omitempty"`
}

func (o CreateAutoCreatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoCreatePolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateAutoCreatePolicyRequest", string(data)}, " ")
}
