package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachSharableVolumesRequest Request Object
type BatchAttachSharableVolumesRequest struct {

	// 共享磁盘ID。
	VolumeId string `json:"volume_id"`

	Body *BatchAttachSharableVolumesRequestBody `json:"body,omitempty"`
}

func (o BatchAttachSharableVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharableVolumesRequest struct{}"
	}

	return strings.Join([]string{"BatchAttachSharableVolumesRequest", string(data)}, " ")
}
