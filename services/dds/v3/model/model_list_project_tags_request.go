package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTagsRequest Request Object
type ListProjectTagsRequest struct {
}

func (o ListProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}
