package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublisherResponse Response Object
type ShowPublisherResponse struct {

	// 发布商详情
	PublisherDetailMap map[string]PublisherVo `json:"publisher_detail_map,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o ShowPublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublisherResponse struct{}"
	}

	return strings.Join([]string{"ShowPublisherResponse", string(data)}, " ")
}
