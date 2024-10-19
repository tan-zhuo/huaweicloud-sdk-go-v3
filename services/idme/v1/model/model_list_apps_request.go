package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppsRequest Request Object
type ListAppsRequest struct {
}

func (o ListAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppsRequest", string(data)}, " ")
}
