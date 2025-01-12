package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTasksResponse Response Object
type DeleteTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTasksResponse struct{}"
	}

	return strings.Join([]string{"DeleteTasksResponse", string(data)}, " ")
}
