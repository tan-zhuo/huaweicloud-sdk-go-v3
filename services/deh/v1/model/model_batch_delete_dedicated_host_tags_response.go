package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDedicatedHostTagsResponse Response Object
type BatchDeleteDedicatedHostTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteDedicatedHostTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDedicatedHostTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDedicatedHostTagsResponse", string(data)}, " ")
}
