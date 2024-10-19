package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerlessPolicyRequest Request Object
type UpdateServerlessPolicyRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateServerlessPolicy `json:"body,omitempty"`
}

func (o UpdateServerlessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerlessPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerlessPolicyRequest", string(data)}, " ")
}
