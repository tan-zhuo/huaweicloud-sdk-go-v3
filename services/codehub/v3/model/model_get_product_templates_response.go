package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetProductTemplatesResponse Response Object
type GetProductTemplatesResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *TemplateListInfo `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetProductTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetProductTemplatesResponse struct{}"
	}

	return strings.Join([]string{"GetProductTemplatesResponse", string(data)}, " ")
}
