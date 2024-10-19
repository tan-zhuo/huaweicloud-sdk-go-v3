package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecretsRequest Request Object
type ListSecretsRequest struct {

	// 每页返回的个数。  默认值：50。
	Limit *string `json:"limit,omitempty"`

	// 分页查询起始的凭据名称，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`

	// 指定事件名称时，仅返回关联该事件的凭据
	EventName *string `json:"event_name,omitempty"`
}

func (o ListSecretsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretsRequest struct{}"
	}

	return strings.Join([]string{"ListSecretsRequest", string(data)}, " ")
}
