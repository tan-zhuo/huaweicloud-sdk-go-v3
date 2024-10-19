package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStreamResponse Response Object
type DeleteStreamResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStreamResponse struct{}"
	}

	return strings.Join([]string{"DeleteStreamResponse", string(data)}, " ")
}
