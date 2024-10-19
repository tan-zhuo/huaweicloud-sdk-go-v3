package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStackSetOperationsResponse Response Object
type ListStackSetOperationsResponse struct {

	// 资源栈集操作列表
	StackSetOperations *[]StackSetOperation `json:"stack_set_operations,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListStackSetOperationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackSetOperationsResponse struct{}"
	}

	return strings.Join([]string{"ListStackSetOperationsResponse", string(data)}, " ")
}
