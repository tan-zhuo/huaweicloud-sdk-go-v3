package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeAccessServiceRequest Request Object
type AuthorizeAccessServiceRequest struct {
	Body *GrantAccessServiceRequestBody `json:"body,omitempty"`
}

func (o AuthorizeAccessServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeAccessServiceRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeAccessServiceRequest", string(data)}, " ")
}
