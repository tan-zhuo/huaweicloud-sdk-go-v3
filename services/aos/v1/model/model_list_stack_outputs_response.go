package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStackOutputsResponse Response Object
type ListStackOutputsResponse struct {

	// 资源栈输出列表
	Outputs        *[]StackOutput `json:"outputs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStackOutputsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackOutputsResponse struct{}"
	}

	return strings.Join([]string{"ListStackOutputsResponse", string(data)}, " ")
}
