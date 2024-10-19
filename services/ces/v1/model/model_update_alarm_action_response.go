package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmActionResponse Response Object
type UpdateAlarmActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionResponse", string(data)}, " ")
}
