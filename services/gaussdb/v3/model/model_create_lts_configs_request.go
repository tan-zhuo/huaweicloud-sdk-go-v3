package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLtsConfigsRequest Request Object
type CreateLtsConfigsRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateLtsConfigs `json:"body,omitempty"`
}

func (o CreateLtsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLtsConfigsRequest struct{}"
	}

	return strings.Join([]string{"CreateLtsConfigsRequest", string(data)}, " ")
}
