package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDomainCertTagsRequest struct {
}

func (o ListDomainCertTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainCertTagsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainCertTagsRequest", string(data)}, " ")
}
