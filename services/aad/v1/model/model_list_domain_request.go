package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainRequest Request Object
type ListDomainRequest struct {
}

func (o ListDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainRequest struct{}"
	}

	return strings.Join([]string{"ListDomainRequest", string(data)}, " ")
}
