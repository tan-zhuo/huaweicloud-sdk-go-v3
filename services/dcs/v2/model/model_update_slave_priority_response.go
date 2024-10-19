package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSlavePriorityResponse Response Object
type UpdateSlavePriorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSlavePriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlavePriorityResponse struct{}"
	}

	return strings.Join([]string{"UpdateSlavePriorityResponse", string(data)}, " ")
}
