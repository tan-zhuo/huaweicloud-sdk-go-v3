package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersForEnterpriseProjectResponse Response Object
type ListUsersForEnterpriseProjectResponse struct {

	// 用户信息。
	Users          *[]ListUsersForEnterpriseProjectResUsers `json:"users,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListUsersForEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersForEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"ListUsersForEnterpriseProjectResponse", string(data)}, " ")
}
