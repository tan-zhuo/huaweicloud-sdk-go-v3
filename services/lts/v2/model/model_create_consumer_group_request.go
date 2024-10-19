package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConsumerGroupRequest Request Object
type CreateConsumerGroupRequest struct {

	// 日志组ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID。 缺省值：None 最小长度：36 最大长度：36
	GroupId string `json:"group_id"`

	// 日志流ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID 缺省值：None 最小长度：36 最大长度：36
	StreamId string `json:"stream_id"`

	Body *ConsumerGroupInfo `json:"body,omitempty"`
}

func (o CreateConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupRequest", string(data)}, " ")
}
