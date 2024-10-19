package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllAppRequest Request Object
type ListAllAppRequest struct {
	Body *ListAllAppRequestBody `json:"body,omitempty"`
}

func (o ListAllAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllAppRequest struct{}"
	}

	return strings.Join([]string{"ListAllAppRequest", string(data)}, " ")
}
