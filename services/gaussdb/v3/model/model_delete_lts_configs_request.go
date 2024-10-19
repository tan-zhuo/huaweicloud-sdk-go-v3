package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLtsConfigsRequest Request Object
type DeleteLtsConfigsRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *DeleteLtsConfigsRequestBody `json:"body,omitempty"`
}

func (o DeleteLtsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsConfigsRequest struct{}"
	}

	return strings.Join([]string{"DeleteLtsConfigsRequest", string(data)}, " ")
}
