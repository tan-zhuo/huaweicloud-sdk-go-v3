package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBindingResponse Response Object
type DeleteBindingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBindingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBindingResponse struct{}"
	}

	return strings.Join([]string{"DeleteBindingResponse", string(data)}, " ")
}
