package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServerTagsResponse Response Object
type BatchDeleteServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerTagsResponse", string(data)}, " ")
}
