package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePortRequestBody
type CreatePortRequestBody struct {
	Port *CreatePortOption `json:"port"`
}

func (o CreatePortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePortRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePortRequestBody", string(data)}, " ")
}
