package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDetailsConsumerGroupRequest Request Object
type ListDetailsConsumerGroupRequest struct {

	// 日志组ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID。 缺省值：None 最小长度：36 最大长度：36
	GroupId string `json:"group_id"`

	// 日志流ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID 缺省值：None 最小长度：36 最大长度：36
	StreamId string `json:"stream_id"`

	// 消费组名称
	ConsumerGroupName string `json:"consumer_group_name"`
}

func (o ListDetailsConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDetailsConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"ListDetailsConsumerGroupRequest", string(data)}, " ")
}
