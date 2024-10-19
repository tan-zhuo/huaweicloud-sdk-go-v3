package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagsResponse Response Object
type CreateTagsResponse struct {

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateTagsResponse", string(data)}, " ")
}
