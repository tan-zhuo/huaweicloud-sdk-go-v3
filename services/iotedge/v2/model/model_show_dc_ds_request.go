package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDcDsRequest Request Object
type ShowDcDsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`
}

func (o ShowDcDsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDcDsRequest struct{}"
	}

	return strings.Join([]string{"ShowDcDsRequest", string(data)}, " ")
}
