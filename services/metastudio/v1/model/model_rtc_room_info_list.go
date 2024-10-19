package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RtcRoomInfoList RTC房间信息。
type RtcRoomInfoList struct {

	// RTC应用ID。
	AppId *string `json:"app_id,omitempty"`

	// RTC房间ID。
	RoomId *string `json:"room_id,omitempty"`

	// 加入RTC房间用户信息。
	Users *[]RtcUserInfo `json:"users,omitempty"`
}

func (o RtcRoomInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcRoomInfoList struct{}"
	}

	return strings.Join([]string{"RtcRoomInfoList", string(data)}, " ")
}
