package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadbalancerTagsResponse Response Object
type ShowLoadbalancerTagsResponse struct {

	// 标签列表
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerTagsResponse", string(data)}, " ")
}
