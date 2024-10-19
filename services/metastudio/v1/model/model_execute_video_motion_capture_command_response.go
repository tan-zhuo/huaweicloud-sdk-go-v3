package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteVideoMotionCaptureCommandResponse Response Object
type ExecuteVideoMotionCaptureCommandResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteVideoMotionCaptureCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteVideoMotionCaptureCommandResponse struct{}"
	}

	return strings.Join([]string{"ExecuteVideoMotionCaptureCommandResponse", string(data)}, " ")
}
