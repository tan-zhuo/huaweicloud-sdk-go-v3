package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskResponse Response Object
type UpdateTaskResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskResponse", string(data)}, " ")
}
