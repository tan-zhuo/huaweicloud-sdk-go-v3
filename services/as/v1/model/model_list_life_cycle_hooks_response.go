package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLifeCycleHooksResponse Response Object
type ListLifeCycleHooksResponse struct {

	// 生命周期挂钩列表。
	LifecycleHooks *[]LifecycleHookList `json:"lifecycle_hooks,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListLifeCycleHooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLifeCycleHooksResponse struct{}"
	}

	return strings.Join([]string{"ListLifeCycleHooksResponse", string(data)}, " ")
}
