package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScriptResponse Response Object
type DeleteScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScriptResponse struct{}"
	}

	return strings.Join([]string{"DeleteScriptResponse", string(data)}, " ")
}
