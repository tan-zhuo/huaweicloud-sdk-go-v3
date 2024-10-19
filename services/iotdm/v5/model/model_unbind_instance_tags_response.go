package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindInstanceTagsResponse Response Object
type UnbindInstanceTagsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnbindInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"UnbindInstanceTagsResponse", string(data)}, " ")
}
