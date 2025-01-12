package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCnfResponse Response Object
type CreateCnfResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateCnfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCnfResponse struct{}"
	}

	return strings.Join([]string{"CreateCnfResponse", string(data)}, " ")
}
