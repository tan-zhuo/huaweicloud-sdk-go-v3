package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCaseResponse Response Object
type DeleteCaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteCaseResponse", string(data)}, " ")
}
