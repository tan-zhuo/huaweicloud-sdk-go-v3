package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskGroupResponse Response Object
type UpdateTaskGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskGroupResponse", string(data)}, " ")
}
