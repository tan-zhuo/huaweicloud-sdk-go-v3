package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerNameRequest Request Object
type UpdateServerNameRequest struct {

	// 源端服务器在主机迁移服务中的ID
	SourceId string `json:"source_id"`

	Body *PutSourceServerBody `json:"body,omitempty"`
}

func (o UpdateServerNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerNameRequest", string(data)}, " ")
}
