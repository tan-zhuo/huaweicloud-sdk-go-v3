package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVersionsRequest Request Object
type ListVersionsRequest struct {
}

func (o ListVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListVersionsRequest", string(data)}, " ")
}
