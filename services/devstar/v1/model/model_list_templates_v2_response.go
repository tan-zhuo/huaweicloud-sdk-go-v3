package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplatesV2Response Response Object
type ListTemplatesV2Response struct {

	// 返回模板的数量。
	Count *int32 `json:"count,omitempty"`

	// 返回关联了失效资源的模板数量。
	InvalidCount *int32 `json:"invalid_count,omitempty"`

	// 返回模板的列表。
	Templates      *[]TemplateInfo `json:"templates,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTemplatesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesV2Response struct{}"
	}

	return strings.Join([]string{"ListTemplatesV2Response", string(data)}, " ")
}
