package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskResponse Response Object
type DeleteTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaskResponse", string(data)}, " ")
}
