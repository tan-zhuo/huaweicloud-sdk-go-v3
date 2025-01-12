package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMediaProcessTaskResponse Response Object
type DeleteMediaProcessTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMediaProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMediaProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteMediaProcessTaskResponse", string(data)}, " ")
}
