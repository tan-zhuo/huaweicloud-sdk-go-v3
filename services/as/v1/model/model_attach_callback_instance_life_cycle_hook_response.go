package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachCallbackInstanceLifeCycleHookResponse Response Object
type AttachCallbackInstanceLifeCycleHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachCallbackInstanceLifeCycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachCallbackInstanceLifeCycleHookResponse struct{}"
	}

	return strings.Join([]string{"AttachCallbackInstanceLifeCycleHookResponse", string(data)}, " ")
}
