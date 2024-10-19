package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEndpointPermissionsResponse Response Object
type AddEndpointPermissionsResponse struct {

	// 白名单记录列表。每个白名单记录的格式为iam:domain::授权账号ID。  其中，授权账号ID是长度为32的字符串，只包含英文字母（a-f）或数字；也可为*，表示允许全部用户连接。
	Permissions *[]string `json:"permissions,omitempty"`

	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddEndpointPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEndpointPermissionsResponse struct{}"
	}

	return strings.Join([]string{"AddEndpointPermissionsResponse", string(data)}, " ")
}
