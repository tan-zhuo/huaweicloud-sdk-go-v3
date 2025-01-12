package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleActionResponse Response Object
type CreateRuleActionResponse struct {

	// 规则动作ID，用于唯一标识一条规则动作，在创建规则动作时由物联网平台分配获得，创建时无需携带，由平台统一分配唯一的action_id。
	ActionId *string `json:"action_id,omitempty"`

	// 规则动作对应的的规则触发条件ID。
	RuleId *string `json:"rule_id,omitempty"`

	// 资源空间ID。
	AppId *string `json:"app_id,omitempty"`

	// 规则动作的类型，取值范围： - HTTP_FORWARDING：HTTP服务消息类型。 - DIS_FORWARDING：转发DIS服务消息类型。 - OBS_FORWARDING：转发OBS服务消息类型。 - AMQP_FORWARDING：转发AMQP服务消息类型。 - DMS_KAFKA_FORWARDING：转发kafka消息类型。 [- ROMA_FORWARDING: 转发Roma消息类型。（仅企业版支持）](tag:hws) - INFLUXDB_FORWARDING：转发InfluxDB消息类型。[（仅标准版和企业版支持）](tag:hws) - MYSQL_FORWARDING：转发MySQL消息类型。[（仅标准版和企业版支持）](tag:hws) - FUNCTIONGRAPH_FORWARDING：转发FunctionGraph消息类型。[（仅标准版和企业版支持）](tag:hws) [- MRS_KAFKA_FORWARDING：转发MRS_KAFKA消息类型。（仅企业版支持）](tag:hws) [- DMS_ROCKETMQ_FORWARDING：转发RocketMQ消息类型。（仅标准版和企业版支持）](tag:hws)
	Channel *string `json:"channel,omitempty"`

	ChannelDetail  *ChannelDetail `json:"channel_detail,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateRuleActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleActionResponse struct{}"
	}

	return strings.Join([]string{"CreateRuleActionResponse", string(data)}, " ")
}
