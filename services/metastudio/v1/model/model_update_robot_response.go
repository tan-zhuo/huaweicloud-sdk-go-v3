package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRobotResponse Response Object
type UpdateRobotResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRobotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRobotResponse struct{}"
	}

	return strings.Join([]string{"UpdateRobotResponse", string(data)}, " ")
}
