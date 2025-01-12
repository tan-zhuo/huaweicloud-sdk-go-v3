package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRequestPropertyResponse Response Object
type DeleteRequestPropertyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRequestPropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRequestPropertyResponse struct{}"
	}

	return strings.Join([]string{"DeleteRequestPropertyResponse", string(data)}, " ")
}
