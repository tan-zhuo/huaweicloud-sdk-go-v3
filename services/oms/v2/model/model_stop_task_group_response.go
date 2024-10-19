package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopTaskGroupResponse Response Object
type StopTaskGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopTaskGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopTaskGroupResponse struct{}"
	}

	return strings.Join([]string{"StopTaskGroupResponse", string(data)}, " ")
}
