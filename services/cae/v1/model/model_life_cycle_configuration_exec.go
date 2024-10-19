package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LifeCycleConfigurationExec 执行命令。
type LifeCycleConfigurationExec struct {

	// shell语句。
	Command *[]string `json:"command,omitempty"`
}

func (o LifeCycleConfigurationExec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifeCycleConfigurationExec struct{}"
	}

	return strings.Join([]string{"LifeCycleConfigurationExec", string(data)}, " ")
}
