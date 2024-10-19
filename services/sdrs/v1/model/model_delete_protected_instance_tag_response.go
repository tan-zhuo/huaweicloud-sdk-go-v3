package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectedInstanceTagResponse Response Object
type DeleteProtectedInstanceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProtectedInstanceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceTagResponse", string(data)}, " ")
}
