package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddSharedTagsResponse Response Object
type BatchAddSharedTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddSharedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddSharedTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddSharedTagsResponse", string(data)}, " ")
}
