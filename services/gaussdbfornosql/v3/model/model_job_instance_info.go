package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobInstanceInfo 执行任务的实例信息。
type JobInstanceInfo struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`
}

func (o JobInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInstanceInfo struct{}"
	}

	return strings.Join([]string{"JobInstanceInfo", string(data)}, " ")
}
