package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveUsersRequest Request Object
type RemoveUsersRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 应用id
	AppId string `json:"app_id"`

	// 房间id
	RoomId string `json:"room_id"`

	Body *RemoveUsersReq `json:"body,omitempty"`
}

func (o RemoveUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveUsersRequest struct{}"
	}

	return strings.Join([]string{"RemoveUsersRequest", string(data)}, " ")
}
