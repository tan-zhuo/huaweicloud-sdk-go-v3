package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerlessPolicyResponse Response Object
type UpdateServerlessPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerlessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerlessPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerlessPolicyResponse", string(data)}, " ")
}
