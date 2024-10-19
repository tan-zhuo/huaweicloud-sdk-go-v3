package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTagsResponse Response Object
type ShowInstanceTagsResponse struct {

	// 标签列表。
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceTagsResponse", string(data)}, " ")
}
