package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthorizableTicketsExternalRequest Request Object
type ListAuthorizableTicketsExternalRequest struct {
	Body *ListAuthorizableTicketsReq `json:"body,omitempty"`
}

func (o ListAuthorizableTicketsExternalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizableTicketsExternalRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizableTicketsExternalRequest", string(data)}, " ")
}
