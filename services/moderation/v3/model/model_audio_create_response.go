package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AudioCreateResponse struct {

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存。
	RequestId *string `json:"request_id,omitempty"`

	// 作业唯一标识。
	JobId *string `json:"job_id,omitempty"`
}

func (o AudioCreateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioCreateResponse struct{}"
	}

	return strings.Join([]string{"AudioCreateResponse", string(data)}, " ")
}
