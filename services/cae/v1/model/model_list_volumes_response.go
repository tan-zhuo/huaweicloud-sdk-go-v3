package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumesResponse Response Object
type ListVolumesResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *VolumeKindObj `json:"kind,omitempty"`

	// 云存储列表。
	Items          *[]Volume `json:"items,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesResponse", string(data)}, " ")
}
