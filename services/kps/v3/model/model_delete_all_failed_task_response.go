package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAllFailedTaskResponse Response Object
type DeleteAllFailedTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAllFailedTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAllFailedTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteAllFailedTaskResponse", string(data)}, " ")
}
