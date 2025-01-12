package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAssetRequest Request Object
type ExportAssetRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *AssetExportRequest `json:"body,omitempty"`
}

func (o ExportAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAssetRequest struct{}"
	}

	return strings.Join([]string{"ExportAssetRequest", string(data)}, " ")
}
