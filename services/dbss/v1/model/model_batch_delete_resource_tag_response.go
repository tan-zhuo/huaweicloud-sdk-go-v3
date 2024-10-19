package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTagResponse Response Object
type BatchDeleteResourceTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagResponse", string(data)}, " ")
}
