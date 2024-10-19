package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateNatTagResponse Response Object
type CreatePrivateNatTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePrivateNatTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatTagResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatTagResponse", string(data)}, " ")
}
