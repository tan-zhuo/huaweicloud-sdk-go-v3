package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBuildJobRequest Request Object
type DeleteBuildJobRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`
}

func (o DeleteBuildJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBuildJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteBuildJobRequest", string(data)}, " ")
}
