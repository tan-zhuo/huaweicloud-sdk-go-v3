package model

import (
	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// KafkaConfigResponseDto **参数说明**：Kafka配置信息。
type KafkaConfigResponseDto struct {

	// **参数说明**：每一套Kafka配置的唯一ID。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）的组合。
	KafkaConfigId *string `json:"kafka_config_id,omitempty"`

	// **参数说明**：kafka的主题列表。  **取值范围**：  - v2x-v1-tracks：edge上报的车辆轨迹数据  - v2x-v1-bsm：车载T-BOX，mqtt协议接入rsu， websocket协议接入rsu上报的BSM消息数据  - v2x-v1-rsi：mqtt协议接入rsu，edge上报的RSI消息数据  - v2x-v1-rsm： mqtt协议接入rsu，edge上报的RSM消息数据  - v2x-v1-spat：mqtt协议接入rsu， websocket协议接入rsu上报的SPAT消息数据  - v2x-v1-edge-flow：edge上报的车流量统计信息数据
	KafkaTopics *[]string `json:"kafka_topics,omitempty"`

	// **参数说明**：Kafka broker列表。
	Brokers *[]string `json:"brokers,omitempty"`

	// **参数说明**：kafka用户名。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）的组合。
	Username *string `json:"username,omitempty"`

	// **参数说明**：Topic前缀。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）的组合。
	TopicPrefix *string `json:"topic_prefix,omitempty"`

	// **参数说明**：一套kafka的连接状态。  **取值范围**：  - OFFLINE：离线  - ONLINE：在线
	Status *KafkaConfigResponseDtoStatus `json:"status,omitempty"`

	// **参数说明**：创建时间。 格式为yyyy-MM-dd'T'HH:mm:ss'Z' 例如：2015-12-12T12:12:12Z
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：修改时间。 格式为yyyy-MM-dd'T'HH:mm:ss'Z' 例如：2015-12-12T12:12:12Z
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`
}

func (o KafkaConfigResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaConfigResponseDto struct{}"
	}

	return strings.Join([]string{"KafkaConfigResponseDto", string(data)}, " ")
}

type KafkaConfigResponseDtoStatus struct {
	value string
}

type KafkaConfigResponseDtoStatusEnum struct {
	OFFLINE KafkaConfigResponseDtoStatus
	ONLINE  KafkaConfigResponseDtoStatus
}

func GetKafkaConfigResponseDtoStatusEnum() KafkaConfigResponseDtoStatusEnum {
	return KafkaConfigResponseDtoStatusEnum{
		OFFLINE: KafkaConfigResponseDtoStatus{
			value: "OFFLINE",
		},
		ONLINE: KafkaConfigResponseDtoStatus{
			value: "ONLINE",
		},
	}
}

func (c KafkaConfigResponseDtoStatus) Value() string {
	return c.value
}

func (c KafkaConfigResponseDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KafkaConfigResponseDtoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
