package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMeetingResponse Response Object
type UpdateMeetingResponse struct {

	// 会议信息列表。
	Body           *[]ConferenceInfo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMeetingResponse struct{}"
	}

	return strings.Join([]string{"UpdateMeetingResponse", string(data)}, " ")
}
