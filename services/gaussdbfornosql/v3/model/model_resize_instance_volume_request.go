package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeInstanceVolumeRequest Request Object
type ResizeInstanceVolumeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResizeInstanceVolumeRequestBody `json:"body,omitempty"`
}

func (o ResizeInstanceVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceVolumeRequest struct{}"
	}

	return strings.Join([]string{"ResizeInstanceVolumeRequest", string(data)}, " ")
}
