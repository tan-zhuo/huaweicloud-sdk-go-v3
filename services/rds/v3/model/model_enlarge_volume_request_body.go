package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnlargeVolumeRequestBody struct {
	EnlargeVolume *EnlargeVolumeObject `json:"enlarge_volume"`
}

func (o EnlargeVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"EnlargeVolumeRequestBody", string(data)}, " ")
}
