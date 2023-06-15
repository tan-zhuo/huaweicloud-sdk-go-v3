package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteCaTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteCaTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCaTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCaTagsResponse", string(data)}, " ")
}
