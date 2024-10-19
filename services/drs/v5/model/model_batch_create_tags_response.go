package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateTagsResponse Response Object
type BatchCreateTagsResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchCreateTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateTagsResponse", string(data)}, " ")
}
