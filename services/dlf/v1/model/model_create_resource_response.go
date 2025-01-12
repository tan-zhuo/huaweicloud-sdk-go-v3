package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceResponse Response Object
type CreateResourceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceResponse", string(data)}, " ")
}
