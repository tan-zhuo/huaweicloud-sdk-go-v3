package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRetentionResponse Response Object
type DeleteRetentionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRetentionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRetentionResponse struct{}"
	}

	return strings.Join([]string{"DeleteRetentionResponse", string(data)}, " ")
}
