package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateListenerTagsResponse Response Object
type BatchCreateListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateListenerTagsResponse", string(data)}, " ")
}
