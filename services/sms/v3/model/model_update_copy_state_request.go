package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCopyStateRequest Request Object
type UpdateCopyStateRequest struct {

	// 源端服务器在主机迁移服务中的ID
	SourceId string `json:"source_id"`

	Body *PutCopyStateReq `json:"body,omitempty"`
}

func (o UpdateCopyStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCopyStateRequest struct{}"
	}

	return strings.Join([]string{"UpdateCopyStateRequest", string(data)}, " ")
}
