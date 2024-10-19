package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRemuxTaskResponse Response Object
type CancelRemuxTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelRemuxTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelRemuxTaskResponse", string(data)}, " ")
}
