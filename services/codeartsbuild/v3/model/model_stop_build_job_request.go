package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBuildJobRequest Request Object
type StopBuildJobRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 构建任务的构建编号，从1开始，每次构建递增1
	BuildNo int32 `json:"build_no"`
}

func (o StopBuildJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBuildJobRequest struct{}"
	}

	return strings.Join([]string{"StopBuildJobRequest", string(data)}, " ")
}
