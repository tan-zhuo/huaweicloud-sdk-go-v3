package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionResponse Response Object
type CreateConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionResponse", string(data)}, " ")
}
