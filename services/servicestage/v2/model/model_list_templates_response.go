package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplatesResponse Response Object
type ListTemplatesResponse struct {

	// 模板列表。
	Templates      *[]TemplateView `json:"templates,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListTemplatesResponse", string(data)}, " ")
}
