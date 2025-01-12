package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDiskInfoRequest Request Object
type UpdateDiskInfoRequest struct {

	// 源端服务器ID
	SourceId string `json:"source_id"`

	Body *PutDiskInfoReq `json:"body,omitempty"`
}

func (o UpdateDiskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDiskInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateDiskInfoRequest", string(data)}, " ")
}
