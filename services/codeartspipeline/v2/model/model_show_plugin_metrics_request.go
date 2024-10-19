package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginMetricsRequest Request Object
type ShowPluginMetricsRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *[]PluginPartQueryDto `json:"body,omitempty"`
}

func (o ShowPluginMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowPluginMetricsRequest", string(data)}, " ")
}
