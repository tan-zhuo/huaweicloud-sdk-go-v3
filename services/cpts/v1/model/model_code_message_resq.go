package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CodeMessageResq struct {

	// 响应码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`
}

func (o CodeMessageResq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeMessageResq struct{}"
	}

	return strings.Join([]string{"CodeMessageResq", string(data)}, " ")
}
