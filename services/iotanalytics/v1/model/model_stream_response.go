package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StreamResponse 流计算
type StreamResponse struct {

	// 输入参数
	Inputs *[]InputResponse `json:"inputs,omitempty"`

	// 流计算任务ID
	JobId *string `json:"job_id,omitempty"`

	// 输出属性，最多支持10个
	Outputs *[]StreamOutput `json:"outputs,omitempty"`
}

func (o StreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamResponse struct{}"
	}

	return strings.Join([]string{"StreamResponse", string(data)}, " ")
}
