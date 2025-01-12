package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStreamForbiddenResponse Response Object
type DeleteStreamForbiddenResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStreamForbiddenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStreamForbiddenResponse struct{}"
	}

	return strings.Join([]string{"DeleteStreamForbiddenResponse", string(data)}, " ")
}
