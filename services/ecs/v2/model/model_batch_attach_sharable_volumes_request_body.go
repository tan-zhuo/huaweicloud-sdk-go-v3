package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachSharableVolumesRequestBody This is a auto create Body Object
type BatchAttachSharableVolumesRequestBody struct {

	// 共享磁盘需要挂载的弹性云服务器列表。
	Serverinfo []BatchAttachSharableVolumesOption `json:"serverinfo"`
}

func (o BatchAttachSharableVolumesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharableVolumesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAttachSharableVolumesRequestBody", string(data)}, " ")
}
