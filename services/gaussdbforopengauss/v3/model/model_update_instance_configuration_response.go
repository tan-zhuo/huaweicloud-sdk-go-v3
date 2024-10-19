package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationResponse Response Object
type UpdateInstanceConfigurationResponse struct {

	// 实例是否需要重启。  - “true”需要重启。 - “false”不需要重启。
	RestartRequired *bool `json:"restart_required,omitempty"`

	// 修改指定实例参数的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationResponse", string(data)}, " ")
}
