package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchEnvTopologyResponse Response Object
type SearchEnvTopologyResponse struct {

	// 组件节点列表。
	NodeList *[]TopoNode `json:"node_list,omitempty"`

	// 组件之间调用指向线列表。
	LineList *[]TopoLine `json:"line_list,omitempty"`

	// 采集器配置。
	CollectorConfig map[string]CollectorConfigModel `json:"collector_config,omitempty"`

	// 开始时间。
	RealStartTime *int64 `json:"real_start_time,omitempty"`

	// 结束时间。
	RealEndTime    *int64 `json:"real_end_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SearchEnvTopologyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchEnvTopologyResponse struct{}"
	}

	return strings.Join([]string{"SearchEnvTopologyResponse", string(data)}, " ")
}
