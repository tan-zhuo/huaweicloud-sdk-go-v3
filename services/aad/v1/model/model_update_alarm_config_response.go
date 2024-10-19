package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmConfigResponse Response Object
type UpdateAlarmConfigResponse struct {

	// SMN的topic urn
	TopicUrn       *string `json:"topic_urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmConfigResponse", string(data)}, " ")
}
