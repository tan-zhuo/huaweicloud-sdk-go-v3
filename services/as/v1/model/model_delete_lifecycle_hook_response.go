package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLifecycleHookResponse Response Object
type DeleteLifecycleHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLifecycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLifecycleHookResponse struct{}"
	}

	return strings.Join([]string{"DeleteLifecycleHookResponse", string(data)}, " ")
}
