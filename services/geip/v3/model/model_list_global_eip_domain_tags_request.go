package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalEipDomainTagsRequest Request Object
type ListGlobalEipDomainTagsRequest struct {
}

func (o ListGlobalEipDomainTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipDomainTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalEipDomainTagsRequest", string(data)}, " ")
}
