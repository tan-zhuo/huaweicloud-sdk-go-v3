package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePartitionResponse Response Object
type UpdatePartitionResponse struct {

	// 资源类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *PartitionMetadata `json:"metadata,omitempty"`

	Spec           *PartitionSpec `json:"spec,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdatePartitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePartitionResponse struct{}"
	}

	return strings.Join([]string{"UpdatePartitionResponse", string(data)}, " ")
}
