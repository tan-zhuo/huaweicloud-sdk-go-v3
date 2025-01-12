package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskOutputForDisplay 作业的输出配置展示
type TaskOutputForDisplay struct {
	Obs *TaskOutputObs `json:"obs,omitempty"`

	Dis *TaskOutputDis `json:"dis,omitempty"`

	Webhook *TaskOutputWebhook `json:"webhook,omitempty"`

	Localpath *TaskOutputLocalpath `json:"localpath,omitempty"`
}

func (o TaskOutputForDisplay) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputForDisplay struct{}"
	}

	return strings.Join([]string{"TaskOutputForDisplay", string(data)}, " ")
}
