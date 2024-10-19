package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSmartChatRoomResponse Response Object
type DeleteSmartChatRoomResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSmartChatRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSmartChatRoomResponse struct{}"
	}

	return strings.Join([]string{"DeleteSmartChatRoomResponse", string(data)}, " ")
}
