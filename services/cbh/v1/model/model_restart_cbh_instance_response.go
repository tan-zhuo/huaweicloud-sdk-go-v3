package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartCbhInstanceResponse Response Object
type RestartCbhInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestartCbhInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCbhInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartCbhInstanceResponse", string(data)}, " ")
}
