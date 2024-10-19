package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDssPoolsRequest Request Object
type ListDssPoolsRequest struct {
}

func (o ListDssPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDssPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListDssPoolsRequest", string(data)}, " ")
}
