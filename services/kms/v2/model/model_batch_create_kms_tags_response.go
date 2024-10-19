package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateKmsTagsResponse Response Object
type BatchCreateKmsTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateKmsTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateKmsTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateKmsTagsResponse", string(data)}, " ")
}
