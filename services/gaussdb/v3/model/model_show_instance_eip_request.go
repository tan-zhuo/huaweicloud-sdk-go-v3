package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceEipRequest Request Object
type ShowInstanceEipRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceEipRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceEipRequest", string(data)}, " ")
}
