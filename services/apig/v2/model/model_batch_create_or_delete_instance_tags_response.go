package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateOrDeleteInstanceTagsResponse Response Object
type BatchCreateOrDeleteInstanceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteInstanceTagsResponse", string(data)}, " ")
}
