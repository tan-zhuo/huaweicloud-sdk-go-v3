package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetBandwidthDomainTagsRequest Request Object
type ListInternetBandwidthDomainTagsRequest struct {
}

func (o ListInternetBandwidthDomainTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthDomainTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthDomainTagsRequest", string(data)}, " ")
}
