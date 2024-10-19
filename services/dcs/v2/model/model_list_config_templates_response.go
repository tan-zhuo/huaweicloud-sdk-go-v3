package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigTemplatesResponse Response Object
type ListConfigTemplatesResponse struct {

	// 模板个数。
	TemplateNum *int32 `json:"template_num,omitempty"`

	// 模板的详情数组。
	Templates      *[]ConfigTemplatesListInfo `json:"templates,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListConfigTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigTemplatesResponse", string(data)}, " ")
}
