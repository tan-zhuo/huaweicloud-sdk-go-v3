package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandInstanceNodesRequest Request Object
type ExpandInstanceNodesRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	Body *EnlargeRequest `json:"body,omitempty"`
}

func (o ExpandInstanceNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodesRequest struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodesRequest", string(data)}, " ")
}
