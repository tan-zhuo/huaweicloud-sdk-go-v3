package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackInstanceResponse Response Object
type RollbackInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RollbackInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackInstanceResponse struct{}"
	}

	return strings.Join([]string{"RollbackInstanceResponse", string(data)}, " ")
}
