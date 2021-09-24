package model

import (
	"encoding/json"

	"strings"
)

// 请求头
type ListEventResponseBodyHeaders struct {
	// 请求长度

	ContentLength *string `json:"content-length,omitempty"`
	// 域名

	Host *string `json:"host,omitempty"`
	// 内容类型

	ContentType *string `json:"content-type,omitempty"`
	// 代理

	UserAgent *string `json:"user-agent,omitempty"`
	// 接收内容类型

	Accept *string `json:"accept,omitempty"`
}

func (o ListEventResponseBodyHeaders) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEventResponseBodyHeaders struct{}"
	}

	return strings.Join([]string{"ListEventResponseBodyHeaders", string(data)}, " ")
}