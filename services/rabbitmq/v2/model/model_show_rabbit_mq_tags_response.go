package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRabbitMqTagsResponse Response Object
type ShowRabbitMqTagsResponse struct {

	// 标签列表
	Tags           *[]TagEntity `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowRabbitMqTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqTagsResponse", string(data)}, " ")
}
