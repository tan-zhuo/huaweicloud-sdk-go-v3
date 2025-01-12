package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubnetsByTagsResponse Response Object
type ListSubnetsByTagsResponse struct {

	// 资源列表
	Resources *[]ListResourceResp `json:"resources,omitempty"`

	// 资源数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubnetsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetsByTagsResponse", string(data)}, " ")
}
