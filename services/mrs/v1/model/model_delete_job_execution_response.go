package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobExecutionResponse Response Object
type DeleteJobExecutionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteJobExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobExecutionResponse struct{}"
	}

	return strings.Join([]string{"DeleteJobExecutionResponse", string(data)}, " ")
}
