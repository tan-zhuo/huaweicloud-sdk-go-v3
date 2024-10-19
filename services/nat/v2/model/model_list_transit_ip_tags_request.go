package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitIpTagsRequest Request Object
type ListTransitIpTagsRequest struct {
}

func (o ListTransitIpTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitIpTagsRequest", string(data)}, " ")
}
