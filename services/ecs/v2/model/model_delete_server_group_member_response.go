package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerGroupMemberResponse Response Object
type DeleteServerGroupMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupMemberResponse", string(data)}, " ")
}
