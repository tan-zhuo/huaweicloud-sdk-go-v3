package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCnfResponse Response Object
type UpdateCnfResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCnfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCnfResponse struct{}"
	}

	return strings.Join([]string{"UpdateCnfResponse", string(data)}, " ")
}
