package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeDomainsRequest Request Object
type AuthorizeDomainsRequest struct {
	Body *AuthorizeDomainsRequestBody `json:"body,omitempty"`
}

func (o AuthorizeDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeDomainsRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeDomainsRequest", string(data)}, " ")
}
