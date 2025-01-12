package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServerGroupMemberResponse Response Object
type AddServerGroupMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddServerGroupMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServerGroupMemberResponse struct{}"
	}

	return strings.Join([]string{"AddServerGroupMemberResponse", string(data)}, " ")
}
