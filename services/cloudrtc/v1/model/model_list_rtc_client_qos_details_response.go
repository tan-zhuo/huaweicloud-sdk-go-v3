package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcClientQosDetailsResponse Response Object
type ListRtcClientQosDetailsResponse struct {

	// 房间ID
	RoomId *string `json:"room_id,omitempty"`

	// QoS质量数据
	Data *[]QosQualityData `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcClientQosDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcClientQosDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListRtcClientQosDetailsResponse", string(data)}, " ")
}
