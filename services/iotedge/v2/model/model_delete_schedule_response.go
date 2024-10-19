package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduleResponse Response Object
type DeleteScheduleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleResponse struct{}"
	}

	return strings.Join([]string{"DeleteScheduleResponse", string(data)}, " ")
}
