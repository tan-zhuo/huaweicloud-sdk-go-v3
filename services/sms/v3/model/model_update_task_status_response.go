package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskStatusResponse Response Object
type UpdateTaskStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResponse", string(data)}, " ")
}
