package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStackSetsResponse Response Object
type ListStackSetsResponse struct {

	// 资源栈集
	StackSets      *[]StackSet `json:"stack_sets,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListStackSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackSetsResponse struct{}"
	}

	return strings.Join([]string{"ListStackSetsResponse", string(data)}, " ")
}
