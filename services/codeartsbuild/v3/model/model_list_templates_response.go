package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplatesResponse Response Object
type ListTemplatesResponse struct {
	Result *QueryTemplatesResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListTemplatesResponse", string(data)}, " ")
}
