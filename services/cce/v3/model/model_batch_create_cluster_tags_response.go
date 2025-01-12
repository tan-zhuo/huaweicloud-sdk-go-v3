package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateClusterTagsResponse Response Object
type BatchCreateClusterTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateClusterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterTagsResponse", string(data)}, " ")
}
