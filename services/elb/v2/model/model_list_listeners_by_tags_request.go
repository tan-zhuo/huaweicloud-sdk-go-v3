package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListListenersByTagsRequest Request Object
type ListListenersByTagsRequest struct {
	Body *ListListenersByTagsRequestBody `json:"body,omitempty"`
}

func (o ListListenersByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsRequest", string(data)}, " ")
}
