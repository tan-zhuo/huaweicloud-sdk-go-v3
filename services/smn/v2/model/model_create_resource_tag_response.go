package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceTagResponse Response Object
type CreateResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceTagResponse", string(data)}, " ")
}
