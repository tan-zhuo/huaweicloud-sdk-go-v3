package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScriptResponse Response Object
type CreateScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScriptResponse struct{}"
	}

	return strings.Join([]string{"CreateScriptResponse", string(data)}, " ")
}
