package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeNodeRequest Request Object
type ShowEdgeNodeRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`
}

func (o ShowEdgeNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeRequest", string(data)}, " ")
}
