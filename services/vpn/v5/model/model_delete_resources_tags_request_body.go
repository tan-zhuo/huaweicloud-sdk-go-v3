package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteResourcesTagsRequestBody struct {
	Tags []ResourceTag `json:"tags"`
}

func (o DeleteResourcesTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourcesTagsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteResourcesTagsRequestBody", string(data)}, " ")
}
