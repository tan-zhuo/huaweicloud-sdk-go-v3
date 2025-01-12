package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaTagsResponse Response Object
type ShowKafkaTagsResponse struct {

	// 标签列表
	Tags           *[]TagEntity `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowKafkaTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaTagsResponse", string(data)}, " ")
}
