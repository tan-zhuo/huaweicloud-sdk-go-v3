package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceConfigsResponse Response Object
type ShowInstanceConfigsResponse struct {

	// kafka配置列表。
	KafkaConfigs   *[]InstanceConfig `json:"kafka_configs,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowInstanceConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigsResponse", string(data)}, " ")
}
