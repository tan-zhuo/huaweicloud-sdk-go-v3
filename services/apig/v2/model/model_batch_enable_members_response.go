package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableMembersResponse Response Object
type BatchEnableMembersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchEnableMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchEnableMembersResponse", string(data)}, " ")
}
