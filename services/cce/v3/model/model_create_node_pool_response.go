package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNodePoolResponse Response Object
type CreateNodePoolResponse struct {

	// API类型，固定值“NodePool”。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v3”。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodePoolMetadata `json:"metadata,omitempty"`

	Spec *NodePoolSpec `json:"spec,omitempty"`

	Status         *CreateNodePoolStatus `json:"status,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CreateNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodePoolResponse struct{}"
	}

	return strings.Join([]string{"CreateNodePoolResponse", string(data)}, " ")
}
