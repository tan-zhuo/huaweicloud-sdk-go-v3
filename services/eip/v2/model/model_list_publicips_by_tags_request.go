package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicipsByTagsRequest Request Object
type ListPublicipsByTagsRequest struct {
	Body *ListPublicipsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListPublicipsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicipsByTagsRequest", string(data)}, " ")
}
