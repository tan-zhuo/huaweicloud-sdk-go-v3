package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectInstanceTagsRequest Request Object
type ListProjectInstanceTagsRequest struct {
}

func (o ListProjectInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectInstanceTagsRequest", string(data)}, " ")
}
