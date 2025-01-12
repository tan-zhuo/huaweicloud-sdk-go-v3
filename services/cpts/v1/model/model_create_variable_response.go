package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVariableResponse Response Object
type CreateVariableResponse struct {

	// 响应码
	Code *string `json:"code,omitempty"`

	Json *CreateVariableResultJson `json:"json,omitempty"`

	// 响应消息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVariableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVariableResponse struct{}"
	}

	return strings.Join([]string{"CreateVariableResponse", string(data)}, " ")
}
