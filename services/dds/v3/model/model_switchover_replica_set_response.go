package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchoverReplicaSetResponse Response Object
type SwitchoverReplicaSetResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchoverReplicaSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverReplicaSetResponse struct{}"
	}

	return strings.Join([]string{"SwitchoverReplicaSetResponse", string(data)}, " ")
}
