package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVolumeTagsRequest Request Object
type ShowVolumeTagsRequest struct {

	// 云硬盘ID
	VolumeId string `json:"volume_id"`
}

func (o ShowVolumeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowVolumeTagsRequest", string(data)}, " ")
}
