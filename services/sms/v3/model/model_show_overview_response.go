package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOverviewResponse Response Object
type ShowOverviewResponse struct {

	// 等待中
	Waiting *int32 `json:"waiting,omitempty"`

	// 复制中
	Replicating *int32 `json:"replicating,omitempty"`

	// 同步中
	Syncing *int32 `json:"syncing,omitempty"`

	// 其它
	Other          *int32 `json:"other,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowOverviewResponse", string(data)}, " ")
}
