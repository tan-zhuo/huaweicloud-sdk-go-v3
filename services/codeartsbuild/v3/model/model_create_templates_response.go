package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplatesResponse Response Object
type CreateTemplatesResponse struct {
	Result *CreateTemplatesItems `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplatesResponse struct{}"
	}

	return strings.Join([]string{"CreateTemplatesResponse", string(data)}, " ")
}
