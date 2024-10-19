package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWelcomeSpeechResponse Response Object
type ShowWelcomeSpeechResponse struct {

	// 欢迎词ID。
	WelcomeSpeechId *string `json:"welcome_speech_id,omitempty"`

	// 欢迎词。
	WelcomeSpeech *string `json:"welcome_speech,omitempty"`

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWelcomeSpeechResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWelcomeSpeechResponse struct{}"
	}

	return strings.Join([]string{"ShowWelcomeSpeechResponse", string(data)}, " ")
}
