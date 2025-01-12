package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubnetTagsResponse Response Object
type ListSubnetTagsResponse struct {

	// tag对象列表
	Tags           *[]ListTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetTagsResponse", string(data)}, " ")
}
