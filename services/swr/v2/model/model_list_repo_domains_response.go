package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepoDomainsResponse Response Object
type ListRepoDomainsResponse struct {

	// 共享租户列表
	Body           *[]ShowRepoDomainsResponse `json:"body,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListRepoDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListRepoDomainsResponse", string(data)}, " ")
}
