package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePublicipTagsResponse Response Object
type BatchDeletePublicipTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeletePublicipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicipTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicipTagsResponse", string(data)}, " ")
}
