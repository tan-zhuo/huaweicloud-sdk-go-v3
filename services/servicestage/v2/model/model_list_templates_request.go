package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplatesRequest Request Object
type ListTemplatesRequest struct {
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}
