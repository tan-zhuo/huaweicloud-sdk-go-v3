package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagResponse Response Object
type CreateTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagResponse struct{}"
	}

	return strings.Join([]string{"CreateTagResponse", string(data)}, " ")
}
