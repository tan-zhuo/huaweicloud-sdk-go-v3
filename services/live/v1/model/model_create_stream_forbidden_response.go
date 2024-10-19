package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStreamForbiddenResponse Response Object
type CreateStreamForbiddenResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateStreamForbiddenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamForbiddenResponse struct{}"
	}

	return strings.Join([]string{"CreateStreamForbiddenResponse", string(data)}, " ")
}
